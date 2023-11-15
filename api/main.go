package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andresromeroh/url-shortener/api/dto"
	"github.com/andresromeroh/url-shortener/api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

var urlService *services.UrlService

func health(c *fiber.Ctx) error {
	response := fiber.Map{
		"status": "ok",
	}
	return c.JSON(response)
}

func shorten(c *fiber.Ctx) error {
	body := new(dto.CreateUrl)
	err := c.BodyParser(body)

	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	createdUrl, err := urlService.CreateUrl(*body)

	if err != nil {
		return err
	}

	return c.JSON(createdUrl)
}

func redirect(c *fiber.Ctx) error {
	url, err := urlService.FindById(
		c.Params("id"),
	)

	if err != nil {
		c.Status(fiber.StatusNotFound).SendString("Not Found")
	}

	return c.Redirect(url.LongURL)
}

func findAllByUserIp(c *fiber.Ctx) error {
	urls, err := urlService.FindAllByUserIp(
		c.Params("ip"),
	)

	if err != nil {
		c.Status(fiber.StatusNotFound).SendString("Not Found")
	}

	return c.JSON(urls)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	urlService = services.NewUrlService()

	app.Get("/:id", redirect)

	v1 := app.Group("/api/v1")

	v1.Get("/health", health)
	v1.Post("/shorten", shorten)
	v1.Get("/:ip/all", findAllByUserIp)

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	log.Fatal(app.Listen(addr))
}
