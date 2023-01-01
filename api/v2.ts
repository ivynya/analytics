import { Router } from "../deps.ts";
import { createPage } from "../notion/createPage.ts";
import { getPage } from "../notion/getPage.ts";
import { queryDatabase } from "../notion/queryDatabase.ts";
import { updateInteractions, updateVisits } from "../notion/updatePage.ts";
import type { Buffer } from "../schema/buffer.ts";

const db = await queryDatabase();
const buffer: Buffer = {};
db.forEach((page) => {
  buffer[page.id] = {
    visits: 0,
    interactions: 0,
  };
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
  .get("/v2/campaign/:id", async (ctx) => {
    const campaign = (await queryDatabase()).find((c) => c.CampaignID === ctx.params.id);
    if (!campaign || campaign.Public.name != "True") {
      ctx.response.status = 400;
      ctx.response.body = "Campaign not found";
    }
    ctx.response.body = JSON.stringify(campaign);
    console.log(`[v2]: GET ${ctx.params.id}`);
  })
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
  .post("/v2/campaign/:id/interaction/:iid", async (ctx) => {
    const campaign = (await queryDatabase()).find((c) => c.CampaignID === ctx.params.id);
    if (!campaign || campaign.Interact.name != "Dynamic") {
      ctx.response.status = 400;
      ctx.response.body = "Cannot create interaction for this campaign"
      return;
    }
    const interactionID = `${ctx.params.id}-${ctx.params.iid}`;
    const interaction = (await queryDatabase()).find((c) => c.CampaignID === interactionID);
    if (interaction) {
      console.log(`[v2]: ${ctx.params.id}::${ctx.params.iid}`);
      buffer[interaction.id].interactions++;
      ctx.response.status = 204;
      return;
    }
    console.log(`[v2]: ${ctx.params.id}++${ctx.params.iid}`);
    const res = await createPage(ctx.params.iid, campaign.CampaignID, campaign.id);
    buffer[interaction.id] = {
      visits: 0,
      interactions: 0,
    };
    if (res.ok) ctx.response.status = 204;
    else ctx.response.status = 400;
  });