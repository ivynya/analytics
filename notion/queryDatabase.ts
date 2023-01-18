
import { flattenQuery } from "../deps.ts";

export async function queryDatabase(): Promise<any[]> {
	const id = Deno.env.get("NOTION_DB_ID");
	const res = await (await fetch(`https://api.notion.com/v1/databases/${id}/query`, {
		method: "POST",
		headers: {
			"Authorization": `Bearer ${Deno.env.get("NOTION_TOKEN")}`,
			"Notion-Version": "2021-08-16"
		}
	})).json();
	
	return flattenQuery(res);
}
