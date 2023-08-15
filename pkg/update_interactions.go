package analytics

import (
	"fmt"

	"github.com/ivynya/analytics/internal/notion"
)

// Update interactions of a campaign, along with any parent campaigns
func UpdateInteractions(nID string, num int) error {
	campaign, err := notion.FetchPage(nID)
	if err != nil {
		return err
	}

	if len(campaign.ParentCampaigns) > 0 {
		for _, parentNID := range campaign.ParentCampaigns {
			err := UpdateInteractions(parentNID, num)
			if err != nil {
				return err
			}
		}
	}

	v := campaign.Interactions + num
	bodyString := fmt.Sprintf(`
		{ "properties": {
			"Interactions": { "number": %d }
		} }`, v)

	err = notion.UpdatePage(nID, bodyString)
	if err != nil {
		return err
	}

	return nil
}
