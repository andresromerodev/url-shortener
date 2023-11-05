package main

import (
	"context"
	"log"
	"os"

	"github.com/andresromeroh/url-shortener/api/db"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type CreateUrlReq struct {
	LongURL string `json:"longUrl"`
}

func getPrismaClient() (*db.PrismaClient, context.Context, error) {
	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		return nil, nil, err
	}

	bg := context.Background()

	return client, bg, nil
}

func health(c *fiber.Ctx) error {
	response := fiber.Map{
		"status": "ok",
	}
	return c.JSON(response)
}

func shorten(c *fiber.Ctx) error {
	request := new(CreateUrlReq)

	if err := c.BodyParser(request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	nanoid, _ := gonanoid.New()
	shortUrl := os.Getenv("BASE_URL") + nanoid

	client, bg, _ := getPrismaClient()

	createdUrl, err := client.URL.CreateOne(
		db.URL.ID.Set(uuid.New().String()),
		db.URL.ShortURL.Set(shortUrl),
		db.URL.LongURL.Set(request.LongURL),
	).Exec(bg)

	if err != nil {
		return err
	}

	return c.JSON(createdUrl)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Get("/api/v1/health", health)
	app.Post("/api/v1/shorten", shorten)

	log.Fatal(app.Listen("127.0.0.1:5000"))
}
