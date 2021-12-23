
import { Application, Router } from "./deps.ts";
import { queryDatabase } from "./notion/queryDatabase.ts";
import { getPage } from "./notion/getPage.ts";
import { updatePage } from "./notion/updatePage.ts";

const app = new Application();
const api = new Router({ prefix: "/api/v1/:id" });

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
	.all("/", async ctx => {
		const campaigns = await queryDatabase();
		const campaign = campaigns.find(c => c.CampaignID === ctx.params.id);
		buffer[campaign.id]++;
		ctx.response.body = campaign;
	});

app.use(api.routes());
app.use(api.allowedMethods());

console.log("[EVT] Listening http://localhost:8000");
await app.listen({ port: 8000 });