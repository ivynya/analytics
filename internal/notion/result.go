package notion

func ConvertPageResult(p PageResult) Page {
	// Get the campaign ID from the result
	campaignID := ""
	for _, richText := range p.Properties.CampaignID.RichText {
		campaignID += richText.PlainText
	}

	// Get the parent campaign IDs from the result
	parentIDs := make([]string, len(p.Properties.ParentCampaign.Relation))
	for i, campaign := range p.Properties.ParentCampaign.Relation {
		parentIDs[i] = campaign.ID
	}

	// Return the formatted page object
	return Page{
		NotionID:        p.ID,
		CampaignID:      campaignID,
		ParentCampaigns: parentIDs,
		Interact:        p.Properties.Interact.Select.Name,
		Public:          p.Properties.Public.Select.Name,
		Visits:          p.Properties.Visits.Number,
		RefVisits:       p.Properties.RefVisits.Number,
		Interactions:    p.Properties.Interactions.Number,
	}
}
