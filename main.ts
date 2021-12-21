
import { config, Application, Router } from "./deps.ts";

const app = new Application();
const router = new Router();
router
	.get("/", ctx => ctx.response.body = "Hello World!")

app.use(router.routes());
app.use(router.allowedMethods());

console.log("[EVT] Listening http://localhost:8000");
await app.listen({ port: 8000 });