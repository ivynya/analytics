# analytics
[![https://maintained.cc/SDBagel/Analytics/1](https://maintained.cc/SDBagel/Analytics/1)](https://ivy.direct/template-analytics)

Notion-integrated simple analytics API.

> ⚠️ This is self-hosted. You're going to need a little bit of server knowledge and a domain if you want this running on the web.

## Setup
1. Duplicate the Notion page and [find the database ID](https://developers.notion.com/docs/working-with-databases)
2. [Create a new Notion integration](https://www.notion.so/my-integrations), copy the API token, and invite it to your duplicated Notion page
3. Clone this repo and create a `.env` file with the ID and token, according to `.env.example`
4. Run `docker build -t analytics .` and `docker run -p 8000:8000 -d analytics`
5. Make a `POST` request to `http://localhost:8000/v1/campaign/promo1` to see it work.

## Manage from Notion

![Notion Template](./.docs/notion_template.png)

The `CampaignID` is what is used in the API. To increment visit (or interaction) count for a campaign, make a `POST` request to `https://localhost:8000/v1/campaign/:CampaignID` (assuming running locally with Docker). Make one post request per increment. This ID can be placed in the appropriate column in an [`SDBagel/redirect`](https://github.com/SDBagel/redirect) instance for tracking.

`RefVisits` tracks referred visits, and is automatically incremented based on sub-campaign. For example, if a `POST` request was made to `/v1/campaign/webpromo`, the visit count for `webpromo` would incremement along with the `Visits` and `RefVisits` count for `promo1`. Specify `ParentCampaign` with the relation column provided.

`Public` determines if the campaign data is returned to the client after making a POST request. If true, the database row is returned as a flattened JSON format. If false, a 204 success is returned.

Additional properties may be added to the Notion table (e.g. a multi-select for better grouping and sorting) and will not affect the analytics. However, modification of existing column names or types will break the API, so be careful.

## Updating Analytics

`analytics` will automatically update when the Notion page is updated. At an absolute maximum, this takes 30 seconds as the database query is cached.

## Invalid Campaigns

Invalid campaign calls will return a 400 non-existent redirect error.
