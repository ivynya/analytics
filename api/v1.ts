
import { Router } from "../deps.ts";
import { getPage } from "../notion/getPage.ts";
import { queryDatabase } from "../notion/queryDatabase.ts";
import { updateVisits } from "../notion/updatePage.ts";

const db = await queryDatabase();
const buffer: { [id: string]: number } = {};
db.forEach(page => {
	buffer[page.id] = 0;
});

setInterval(async () => {
	const queue = Object.keys(buffer).filter(id => buffer[id] > 0);
	const pages = await Promise.all(queue.map(id => getPage(id)));
	await Promise.all(pages.map(page => updateVisits(page, buffer[page.id])));
	Object.keys(buffer).forEach(id => buffer[id] = 0);
}, 5000);

export const api = new Router();
api
	.post("/v1/campaign/:id", async ctx => {
		const campaigns = await queryDatabase();
		const campaign = campaigns.find(c => c.CampaignID === ctx.params.id);
		if (campaign) {
      console.log(`[v1]: ${ctx.params.id}`)
			buffer[campaign.id]++;
			if (campaign.Public.name == "True")
				ctx.response.body = campaign;
			else ctx.response.status = 204;
		} else {
			ctx.response.status = 400;
			ctx.response.body = "Campaign not found";
		}
	});