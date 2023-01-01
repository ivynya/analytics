
export async function createPage(name: string, parentName: string, parentID: string|undefined) {
	return await fetch(`https://api.notion.com/v1/pages`, {
		method: "POST",
		headers: {
			"Authorization": `Bearer ${Deno.env.get("NOTION_TOKEN")}`,
      "Content-Type": "application/json",
			"Notion-Version": "2022-06-28"
		},
    body: JSON.stringify({
      "parent": { "database_id": Deno.env.get("NOTION_DB_ID") },
      "properties": {
        "Campaign": { "title": [{ "text": { "content": name } }] },
        "CampaignID": { "rich_text": [{ "text": { "content": parentName + "-" + name } }] },
        "ParentCampaign": parentID ? { "relation": [{ "id": parentID }] } : undefined,
        "Visits": { "number": 0 },
        "RefVisits": { "number": 0 },
        "Interactions": { "number": 0 }
      }
    })
	});
}