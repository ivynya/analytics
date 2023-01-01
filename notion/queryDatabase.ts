
import { flattenQuery } from "../deps.ts";

const cachePath = `${Deno.cwd()}/notion/cache.json`;
let lastUpdated = new Date(0);

export async function queryDatabase(): Promise<any[]> {
	if (lastUpdated.getTime() + 10000 > Date.now())
		return JSON.parse(await Deno.readTextFile(cachePath));
	else lastUpdated = new Date();
	
	const id = Deno.env.get("NOTION_DB_ID");
	const res = await (await fetch(`https://api.notion.com/v1/databases/${id}/query`, {
		method: "POST",
		headers: {
			"Authorization": `Bearer ${Deno.env.get("NOTION_TOKEN")}`,
			"Notion-Version": "2021-08-16"
		}
	})).json();
	
	const flatRes = flattenQuery(res);
	await Deno.writeTextFile(cachePath, JSON.stringify(flatRes));
	return flatRes;
}
