package analytics

import (
	"fmt"

	"github.com/ivynya/analytics/internal/notion"
)

// Update vists of a campaign, along with any parent campaigns
// If ref is true, will update ref visits of the campaign
func UpdateVisits(nID string, num int, ref bool) error {
	campaign, err := notion.FetchPage(nID)
	if err != nil {
		return err
	}

	if len(campaign.ParentCampaigns) > 0 {
		for _, parentNID := range campaign.ParentCampaigns {
			err := UpdateVisits(parentNID, num, true)
			if err != nil {
				return err
			}
		}
	}

	v := campaign.Visits + num
	r := campaign.RefVisits
	if ref {
		r += num
	}
	bodyString := fmt.Sprintf(`
		{ "properties": {
			"Visits": { "number": %d },
			"RefVisits": { "number": %d }
		} }`, v, r)

	err = notion.UpdatePage(nID, bodyString)
	if err != nil {
		return err
	}

	return nil
}
