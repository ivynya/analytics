package analytics

import (
	"errors"

	"github.com/ivynya/analytics/internal/notion"
)

// Find a campaign by its CampaignID property
func FindCampaignByCID(cID string) (notion.Page, error) {
	db, err := notion.FetchDatabase()
	if err != nil {
		return notion.Page{}, err
	}

	for _, pageResult := range db.Results {
		page := notion.ConvertPageResult(pageResult)
		if page.CampaignID == cID {
			return page, nil
		}
	}

	return notion.Page{}, errors.New("campaign not found")
}

// Find a campaign by its NotionID/raw ID property
func FindCampaignByNID(nID string) (notion.Page, error) {
	db, err := notion.FetchDatabase()
	if err != nil {
		return notion.Page{}, err
	}

	for _, pageResult := range db.Results {
		page := notion.ConvertPageResult(pageResult)
		if page.NotionID == nID {
			return page, nil
		}
	}

	return notion.Page{}, errors.New("campaign not found")
}
