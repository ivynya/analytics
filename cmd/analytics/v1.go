package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	analytics "github.com/ivynya/analytics/pkg"
)

func createRouterV1(a *fiber.App) fiber.Router {
	v1 := a.Group("/v1")

	// Deprecated - will be removed in analytics v4
	// Functionally identical to /v2/campaign/:cID (use instead)
	v1.Post("/campaign/:cID", func(c *fiber.Ctx) error {
		log.Println("[ERR] Deprecated call: /v1/campaign/" + c.Params("cID"))
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

	return v1
}
