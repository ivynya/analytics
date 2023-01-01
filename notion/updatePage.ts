
import { getPage } from "./getPage.ts";

function genVisit(num: number) {
	return { "properties": {
		"Visits": { "number": num } }
	};
}

function genRefVisit(num: number, refNum: number) {
	return {
		"properties": {
			"Visits": { "number": num },
			"RefVisits": { "number": refNum }
		}
	};
}

export function updateInteractions(page: any, num: number) {
	if (page.ParentCampaign?.length > 0) {
		page.ParentCampaign.forEach(async (obj: {id: string}) => {
			const parentPage = await getPage(obj.id);
			await updateInteractions(parentPage, num);
		})
	}

  const i = page.Interactions + num;
  return fetch(`https://api.notion.com/v1/pages/${page.id}`, {
    body: JSON.stringify({
      "properties": {
        "Interactions": { "number": i },
      },
    }),
    method: "PATCH",
    headers: {
      "Authorization": `Bearer ${Deno.env.get("NOTION_TOKEN")}`,
      "Content-Type": "application/json",
      "Notion-Version": "2021-08-16",
    },
  });
}


export function updateVisits(page: any, num: number, ref=false) {
	if (page.ParentCampaign?.length > 0) {
		page.ParentCampaign.forEach(async (obj: {id: string}) => {
			const parentPage = await getPage(obj.id);
			await updateVisits(parentPage, num, true);
		})
	}

	const v = page.Visits + num;
	const rV = page.RefVisits + num;
	return fetch(`https://api.notion.com/v1/pages/${page.id}`, {
		body: JSON.stringify(ref ? genRefVisit(v, rV) : genVisit(v)),
		method: "PATCH",
		headers: {
			"Authorization": `Bearer ${Deno.env.get("NOTION_TOKEN")}`,
			"Content-Type": "application/json",
			"Notion-Version": "2021-08-16"
		}
	});
}