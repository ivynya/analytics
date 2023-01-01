import { Router } from "../deps.ts";
import { getPage } from "../notion/getPage.ts";
import { queryDatabase } from "../notion/queryDatabase.ts";
import { updateInteractions, updateVisits } from "../notion/updatePage.ts";
import type { Buffer } from "../schema/buffer.ts";

const db = await queryDatabase();
const buffer: Buffer = {};
db.forEach((page) => {
  buffer[page.id].visits = 0;
  buffer[page.id].interactions = 0;
});

setInterval(() => {
  Object.keys(buffer)
    .filter((id) => buffer[id].interactions > 0 || buffer[id].visits > 0)
    .map(async (id) => {
      const page = await getPage(id);
      await updateInteractions(page, buffer[id].interactions)
      await updateVisits(page, buffer[id].visits);
      buffer[id].interactions = 0;
      buffer[id].visits = 0;
    });
}, 5000);

export const api = new Router();
api
  .post("/v2/campaign/:id", async (ctx) => {
    const campaign = (await queryDatabase()).find((c) => c.CampaignID === ctx.params.id);
    if (!campaign) {
      ctx.response.status = 400;
      ctx.response.body = "Campaign not found";
    }
    console.log(`[v2]: ${ctx.params.id}`);
    buffer[campaign.id].visits++;
    if (campaign.Public.name == "True")
      ctx.response.body = campaign;
    else ctx.response.status = 204;
  })
  .post("/v2/campaign/:id/interaction", async (ctx) => {
    const campaign = (await queryDatabase()).find((c) => c.CampaignID === ctx.params.id);
    if (!campaign) {
      ctx.response.status = 400;
      ctx.response.body = "Campaign not found";
    }
    console.log(`[v2]: ${ctx.params.id}::${ctx.params.iid}`);
    buffer[campaign.id].interactions++;
    if (campaign.Public.name == "True")
      ctx.response.body = campaign;
    else ctx.response.status = 204;
  })
  .post("/v2/campaign/:id/interaction/:iid", (ctx) => {
    ctx.response.status = 501;
  });