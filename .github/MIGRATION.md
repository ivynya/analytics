
# Migration Guide

This migration guide is a comprehensive guide to the changes you need to make in Notion and in your code to upgrade your analytics version.

## V2 -> V3 Migration

### In Notion
V3 uses the exact same database format as V2, so you can keep your existing Notion database and integration.

### In Code
You need to update any services calling the analytics API.
1. All V2 endpoints have a V3 endpoint that is functionally equivalent, so you can just update the endpoint from /v2/ to /v3/ as shown below.

| method | old endpoint | new endpoint |
|---|---|---|
| GET | `/v2/campaign/:id` | `/v3/campaign/:id` |
| POST | `/v2/campaign/:id` | `/v3/campaign/:id` |
| POST | `/v2/campaign/:id/interaction` | `/v3/campaign/:id/interaction` |
| POST | `/v2/campaign/:id/interaction/:interactionId` | `/v3/campaign/:id/interaction/:interactionId` |
| POST | `/v2/campaign/:id/visit/:interactionId` | `/v3/campaign/:id/visit/:interactionId` |

2. If you use JSON responses from the API (If `Public`=`True` for any of your campaigns), then you need to alter your code to use the new response schema as shown below. If you aren't using Public campaigns, then successes will continue returning 204 status codes and you can ignore this step.

| property name | type |
|---|---|
| `id` | `string` |
| `campaign_id` | `string` |
| `parent_campaigns` | `string[]` |
| `interact` | `string` |
| `public` | `string` |
| `ref_visits` | `number` |
| `visits` | `number` |
| `interactions` | `number` |

You can find the documentation on the rest of the new endpoints in the README.

---

## V1 -> V2 Migration

### In Notion
Using your existing Notion database:
1. Create an property of type `Number` named `Interactions`.
2. Create a property of type `Select` named `Interact` with the options `Enabled`, `Disabled`, and `Dynamic`.
3. Ensure you have a property called `CreatedBy` of type `Created by`.
4. Enable your database's `Subtasks` feature (using the `ParentCampaign` and `SubCampaigns` property names respectively) for better visual organization. Alternatively, just ensure the `ParentCampaign` of type `Relation` property exists.

### In Code
The only API endpoint in V1 is forward-compatible with V2, so you can just update the endpoint from /v1/ to /v2/ as shown below.

| method | old endpoint | new endpoint |
|---|---|---|
| POST | `/v1/campaign/:id` | `/v2/campaign/:id` |

You can find the documentation on the rest of the new endpoints in the README.

Your new Analytics instance and database format is still compatible with your `redirect` instance, if you are running the two alongside each other.