
import { flatten } from "../deps.ts";

export async function getPage(id: any) {
	const res = await (await fetch(`https://api.notion.com/v1/pages/${id}`, {
		method: "GET",
		headers: {
			"Authorization": `Bearer ${Deno.env.get("NOTION_TOKEN")}`,
			"Notion-Version": "2021-08-16"
		}
	})).json();
	
	return flatten(res);
}