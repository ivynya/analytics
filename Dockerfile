
FROM denoland/deno:alpine-1.23.1

EXPOSE 8000

WORKDIR /app

COPY deps.ts .
COPY . .

RUN deno cache deps.ts
RUN deno cache main.ts

CMD ["run", "-A", "main.ts"]