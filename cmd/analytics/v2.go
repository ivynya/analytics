package main

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ivynya/analytics/internal/notion"
	analytics "github.com/ivynya/analytics/pkg"
)

func createRouterV2(a *fiber.App) fiber.Router {
	v2 := a.Group("/v2")

	// Returns a campaign if it is public as PageResult form
	v2.Get("/campaign/:cID", func(c *fiber.Ctx) error {
		cID := c.Params("cID")
		campaign, err := analytics.FindCampaignByCID(cID)
		if err != nil || campaign.Public != "True" {
			return c.Status(400).SendString("Campaign not found")
		}
		unformattedCampaign, err := v2FindCampaignByCIDWithoutConvert(cID)
		if err != nil {
			return c.Status(400).SendString("Campaign not found")
		}
		jso, err := json.Marshal(unformattedCampaign)
		if err != nil {
			return c.Status(500).SendString("Internal server error")
		}
		return c.SendString(string(jso))
	})

	// Increments visits and returns if public as PageResult / 204
	v2.Post("/campaign/:cID", func(c *fiber.Ctx) error {
		cID := c.Params("cID")
		campaign, err := analytics.FindCampaignByCID(cID)
		if err != nil {
			return c.Status(400).SendString("Campaign not found")
		}
		bufferData(campaign.NotionID, 1, 0)

		if campaign.Public == "True" {
			unformattedCampaign, err := v2FindCampaignByCIDWithoutConvert(cID)
			if err != nil {
				return c.Status(400).SendString("Campaign not found")
			}
			jso, err := json.Marshal(unformattedCampaign)
			if err != nil {
				return c.Status(500).SendString("Internal server error")
			}
			return c.SendString(string(jso))
		}

		return c.SendStatus(204)
	})

	// Increments interactions and returns if public as PageResult / 204
	v2.Post("/campaign/:cID/interaction", func(c *fiber.Ctx) error {
		cID := c.Params("cID")
		campaign, err := analytics.FindCampaignByCID(cID)
		if err != nil {
			return c.Status(400).SendString("Campaign not found")
		}
		bufferData(campaign.NotionID, 0, 1)

		if campaign.Public == "True" {
			unformattedCampaign, err := v2FindCampaignByCIDWithoutConvert(cID)
			if err != nil {
				return c.Status(400).SendString("Campaign not found")
			}
			jso, err := json.Marshal(unformattedCampaign)
			if err != nil {
				return c.Status(500).SendString("Internal server error")
			}
			return c.SendString(string(jso))
		}

		return c.SendStatus(204)
	})

	// If parent dynamic, create/increment KPI interaction and returns 204
	v2.Post("/campaign/:cID/interaction/:iID", func(c *fiber.Ctx) error {
		cID := c.Params("cID")
		campaign, err := analytics.FindCampaignByCID(cID)
		if err != nil {
			return c.Status(400).SendString("Campaign not found")
		}
		if campaign.Interact != "Dynamic" {
			return c.Status(400).SendString("Cannot create interaction for campaign")
		}
		iID := c.Params("cID") + "-" + c.Params("iID")
		interaction, err := analytics.FindCampaignByCID(iID)
		if err == nil {
			bufferData(interaction.NotionID, 0, 1)
			return c.SendStatus(204)
		} else {
			err := analytics.CreatePage(iID, cID, campaign.NotionID)
			if err != nil {
				return c.Status(400).SendString("Failed creating KPI")
			}
			return c.SendStatus(204)
		}
	})

	// If parent dynamic, create/increment KPI visit and returns 204
	v2.Post("/campaign/:cID/visit/:iID", func(c *fiber.Ctx) error {
		cID := c.Params("cID")
		campaign, err := analytics.FindCampaignByCID(cID)
		if err != nil {
			return c.Status(400).SendString("Campaign not found")
		}
		if campaign.Interact != "Dynamic" {
			return c.Status(400).SendString("Cannot create interaction for campaign")
		}
		iID := c.Params("cID") + "-" + c.Params("iID")
		interaction, err := analytics.FindCampaignByCID(iID)
		if err == nil {
			bufferData(interaction.NotionID, 1, 0)
			return c.SendStatus(204)
		} else {
			err := analytics.CreatePage(iID, cID, campaign.NotionID)
			if err != nil {
				return c.Status(400).SendString("Failed creating KPI")
			}
			return c.SendStatus(204)
		}
	})

	return v2
}

// This is a private bridge function for v1/v2 API endpoints
// since the v2 API endpoints need to return the unformatted
// notion.PageResult instead of the formatted notion.Campaign
func v2FindCampaignByCIDWithoutConvert(cID string) (notion.PageResult, error) {
	db, err := notion.FetchDatabase()
	if err != nil {
		return notion.PageResult{}, err
	}

	for _, pageResult := range db.Results {
		page := notion.ConvertPageResult(pageResult)
		if page.CampaignID == cID {
			return pageResult, nil
		}
	}

	return notion.PageResult{}, errors.New("campaign not found")
}
