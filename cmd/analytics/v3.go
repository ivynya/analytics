package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/ivynya/analytics/internal/notion"
	analytics "github.com/ivynya/analytics/pkg"
)

func createRouterV3(a *fiber.App) fiber.Router {
	v3 := a.Group("/v3")

	// Returns a campaign if it is public as a Page
	// 0, 0 = visits/interactions (read-only)
	v3.Get("/campaign/:cID", func(c *fiber.Ctx) error {
		return v3UpdateCampaign(c, 0, 0)
	})

	// Increment campaign visits
	// 1, 0 = visits/interactions (+1 visit, 0 interactions)
	v3.Post("/campaign/:cID", func(c *fiber.Ctx) error {
		return v3UpdateCampaign(c, 1, 0)
	})

	// Increment campaign interactions
	// 0, 1 = visits/interactions (0 visits, +1 interaction)
	v3.Post("/campaign/:cID/interaction", func(c *fiber.Ctx) error {
		return v3UpdateCampaign(c, 0, 1)
	})

	// Returns an KPI if it is public as a Page
	// 0, 0 = visits/interactions (read-only)
	v3.Get("/campaign/:cID/interaction/:iID", func(c *fiber.Ctx) error {
		return v3UpdateInteraction(c, 0, 0)
	})

	// Create/increment KPI interaction
	// 0, 1 = visits/interactions (0 visits, +1 interaction)
	v3.Post("/campaign/:cID/interaction/:iID", func(c *fiber.Ctx) error {
		return v3UpdateInteraction(c, 0, 1)
	})

	// Create/increment KPI visit
	// 1, 0 = visits/interactions (+1 visit, 0 interactions)
	v3.Post("/campaign/:cID/visit/:iID", func(c *fiber.Ctx) error {
		return v3UpdateInteraction(c, 1, 0)
	})

	return v3
}

// If the campaign exists, buffer the data and return v3ResponseProtocol
// On error (campaign does not exist), return 400
func v3UpdateCampaign(c *fiber.Ctx, visits int, interactions int) error {
	campaign, err := analytics.FindCampaignByCID(c.Params("cID"))
	if err != nil {
		return c.Status(400).SendString("Campaign not found")
	}
	bufferData(campaign.NotionID, visits, interactions)
	return v3ResponseProtocol(c, campaign)
}

// If the interaction exists, buffer the data and return v3ResponseProtocol
// If the interaction does not exist, create it and return 204
// On error (campaign does not exist, is not Dynamic), return 400
func v3UpdateInteraction(c *fiber.Ctx, visits int, interactions int) error {
	cID := c.Params("cID")
	campaign, err := analytics.FindCampaignByCID(cID)
	if err != nil {
		return c.Status(400).SendString("Campaign does not exist")
	}
	if campaign.Interact == "Disabled" {
		return c.Status(400).SendString("Campaign interactions disabled")
	}
	iID := c.Params("cID") + "-" + c.Params("iID")
	interaction, err := analytics.FindCampaignByCID(iID)
	if err == nil {
		bufferData(interaction.NotionID, visits, interactions)
		return v3ResponseProtocol(c, campaign)
	} else {
		if campaign.Interact != "Dynamic" {
			return c.Status(400).SendString("Campaign dynamic interactions disabled")
		}
		err := analytics.CreatePage(iID, cID, campaign.NotionID)
		if err != nil {
			return c.Status(400).SendString("Failed creating KPI")
		}
		return c.SendStatus(204)
	}
}

// If the campaign is public, return the campaign as JSON (Page)
// If the campaign is not public, return 204
func v3ResponseProtocol(c *fiber.Ctx, campaign notion.Page) error {
	if campaign.Public == "True" {
		jso, err := json.Marshal(campaign)
		if err != nil {
			return c.Status(500).SendString("Internal server error")
		}
		return c.SendString(string(jso))
	} else {
		return c.SendStatus(204)
	}
}
