package analytics

import (
	"fmt"
	"os"

	"github.com/ivynya/analytics/internal/notion"
)

// Create a new campaign page in Notion given CampaignID
// and ParentCampaignID with parent campaign's NotionID
func CreatePage(cID string, parentCID string, parentNID string) error {
	body := fmt.Sprintf(`
		{
			"parent": { "database_id": "%s" },
			"properties": {
				"Campaign": { "title": [{ "text": { "content": %s } }] },
				"CampaignID": { "rich_text": [{ "text": { "content": %s-%s } }] },
				"ParentCampaign": { "relation": [{ "id": { "content": "%s" } }] },
				"Visits": { "number": 0 },
				"RefVisits": { "number": 0 },
				"Interactions": { "number": 0 }
			}
		}
	`, os.Getenv("NOTION_DB_ID"), cID, parentCID, cID, parentNID)

	return notion.CreatePage(body)
}
