package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	go bufferFlushLoop()

	app := fiber.New()
	app.Use(cors.New())

	log.Println("[analytics] creating routers...")

	createRouterV1(app)
	createRouterV2(app)
	createRouterV3(app)

	tokenPresent := os.Getenv("NOTION_TOKEN") != ""
	log.Println("[analytics] $NOTION_DB_ID:", os.Getenv("NOTION_DB_ID"))
	log.Println("[analytics] $NOTION_TOKEN:", tokenPresent)
	log.Println("[analytics] starting fiber on port 3000")

	log.Fatal(app.Listen(":3000"))
}
