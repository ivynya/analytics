package notion

type DatabaseResult struct {
	Results    []PageResult `json:"results"`
	NextCursor string       `json:"next_cursor"`
	HasMore    bool         `json:"has_more"`
}

type PageResult struct {
	ID         string `json:"id"`
	Properties struct {
		CampaignID struct {
			RichText []struct {
				PlainText string `json:"plain_text"`
			} `json:"rich_text"`
		} `json:"CampaignID"`
		RefVisits struct {
			Number int `json:"number"`
		} `json:"RefVisits"`
		Visits struct {
			Number int `json:"number"`
		} `json:"Visits"`
		Interactions struct {
			Number int `json:"number"`
		} `json:"Interactions"`
		ParentCampaign struct {
			Relation []struct {
				ID string `json:"id"`
			} `json:"relation"`
			HasMore bool `json:"has_more"`
		} `json:"ParentCampaign"`
		Interact struct {
			Select struct {
				Name string `json:"name"`
			} `json:"select"`
		} `json:"Interact"`
		Public struct {
			Select struct {
				Name string `json:"name"`
			} `json:"select"`
		} `json:"Public"`
	} `json:"properties"`
}

type Page struct {
	NotionID        string   `json:"id"`
	CampaignID      string   `json:"campaign_id"`
	ParentCampaigns []string `json:"parent_campaigns"`
	Interact        string   `json:"interact"`
	Public          string   `json:"public"`
	Visits          int      `json:"visits"`
	RefVisits       int      `json:"ref_visits"`
	Interactions    int      `json:"interactions"`
}
