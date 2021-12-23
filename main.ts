
import { Application, Router } from "./deps.ts";
import { queryDatabase } from "./notion/queryDatabase.ts";
import { getPage } from "./notion/getPage.ts";
import { updatePage } from "./notion/updatePage.ts";

const app = new Application();
const api = new Router({ prefix: "/v1/campaign/:id" });

const db = await queryDatabase();
const buffer: { [id: string]: number } = {};
db.forEach(page => {
	buffer[page.id] = 0;
});

setInterval(async () => {
	const queue = Object.keys(buffer).filter(id => buffer[id] > 0);
	const pages = await Promise.all(queue.map(id => getPage(id)));
	await Promise.all(pages.map(page => updatePage(page, buffer[page.id])));
	Object.keys(buffer).forEach(id => buffer[id] = 0);
}, 5000);

api
	.post("/", async ctx => {
		const campaigns = await queryDatabase();
		const campaign = campaigns.find(c => c.CampaignID === ctx.params.id);
		if (campaign) {
			console.log(`[LOG] Post: ${campaign.CampaignID}`);
			buffer[campaign.id]++;
			if (campaign.Public.name == "True")
				ctx.response.body = campaign;
			else ctx.response.status = 204;
		} else {
			ctx.response.status = 404;
			ctx.response.body = "Not found";
		}
	});

app.use(api.routes());
app.use(api.allowedMethods());

console.log("[EVT] Listening http://localhost:8000");
await app.listen({ port: 8000 });