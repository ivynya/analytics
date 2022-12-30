
import { Router } from "../deps.ts";
import { getPage } from "../notion/getPage.ts";
import { queryDatabase } from "../notion/queryDatabase.ts";

const db = await queryDatabase();
const buffer: { [id: string]: number } = {};
db.forEach(page => {
	buffer[page.id] = 0;
});

setInterval(async () => {
	const queue = Object.keys(buffer).filter(id => buffer[id] > 0);
	const pages = await Promise.all(queue.map(id => getPage(id)));
	await Promise.all(pages.map(page => updateInteractions(page, buffer[page.id])));
	Object.keys(buffer).forEach(id => buffer[id] = 0);
}, 5000);

export const api = new Router();
api
	.post("/v2/campaign/:id/interaction", async ctx => {
		const campaigns = await queryDatabase();
		const campaign = campaigns.find(c => c.CampaignID === ctx.params.id);
		if (campaign) {
      console.log(`[v2]: ${ctx.params.id}::${ctx.params.iid}`)
			buffer[campaign.id]++;
			if (campaign.Public.name == "True")
				ctx.response.body = campaign;
			else ctx.response.status = 204;
		} else {
			ctx.response.status = 400;
			ctx.response.body = "Campaign not found";
		}
	});

function updateInteractions(page: any, num: number) {
  const i = page.Interactions + num;
  return fetch(`https://api.notion.com/v1/pages/${page.id}`, {
    body: JSON.stringify({ "properties": {
      "Interactions": { "number": i } }
    }),
    method: "PATCH",
    headers: {
      "Authorization": `Bearer ${Deno.env.get("NOTION_TOKEN")}`,
      "Content-Type": "application/json",
      "Notion-Version": "2021-08-16"
    }
  });
}