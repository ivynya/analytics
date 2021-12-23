
function genVisit(num: number) {
	return { "properties": { "Visits": { "number": num } } };
}

function genRefVisit(num: number, refNum: number) {
	return { "properties": {
		"Vists": { "number": num },
		"RefVisits": { "number": refNum }
	} };
}

export async function updatePage(page: any, num: number, ref = false) {
	console.log(`[LOG] Updating ${page.id} : ${num} ${ref?': Ref':''}`);
	const v = page.Visits + num;
	const rV = page.RefVisits + num;
	return fetch(`https://api.notion.com/v1/pages/${page.id}`, {
		body: JSON.stringify(ref ? genRefVisit(v, rV) : genVisit(v)),
		method: "PATCH",
		headers: {
			"Authorization": `Bearer ${Deno.env.get("NOTION_TOKEN")}`,
			"Content-Type": "application/json",
			"Notion-Version": "2021-08-16"
		}
	});
}