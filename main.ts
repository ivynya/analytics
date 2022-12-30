
import { Application } from "./deps.ts";
import { api as api_v1 } from "./api/v1.ts";
import { api as api_v2 } from "./api/v2.ts";

const app = new Application();

app.use(api_v1.routes());
app.use(api_v1.allowedMethods());
app.use(api_v2.routes());
app.use(api_v2.allowedMethods());

console.log("[EVT] Listening http://localhost:8000");
await app.listen({ port: 8000 });