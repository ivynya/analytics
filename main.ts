
import { Application, Router } from "./deps.ts";

const app = new Application();
const api = new Router({ prefix: "/api/v1/:id" });

api
	.get("/", ctx => ctx.response.body = ctx.params.id)
	.post("/", ctx => ctx.response.body = ctx.params.id);

app.use(api.routes());
app.use(api.allowedMethods());

console.log("[EVT] Listening http://localhost:8000");
await app.listen({ port: 8000 });