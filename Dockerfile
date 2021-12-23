
FROM denoland/deno:alpine-1.17.0

EXPOSE 8000

WORKDIR /app

COPY deps.ts .
COPY . .

RUN deno cache deps.ts
RUN deno cache main.ts

CMD ["run", "-A", "main.ts"]