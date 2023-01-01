
import { flattenPage } from "../deps.ts";

export async function getPage(id: string) {
	const res = await (await fetch(`https://api.notion.com/v1/pages/${id}`, {
		method: "GET",
		headers: {
			"Authorization": `Bearer ${Deno.env.get("NOTION_TOKEN")}`,
			"Notion-Version": "2021-08-16"
		}
	})).json();
	
	return flattenPage(res);
}