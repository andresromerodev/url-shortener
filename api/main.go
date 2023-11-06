package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andresromeroh/url-shortener/api/dto"
	"github.com/andresromeroh/url-shortener/api/services"
	"github.com/gofiber/fiber/v2"
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

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	urlService = services.NewUrlService()

	// To make it shorter exclude the redirect from the api/v1 group
	app.Get("/:id", redirect)

	v1 := app.Group("/api/v1")

	v1.Get("/health", health)
	v1.Post("/shorten", shorten)

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	log.Fatal(app.Listen(addr))
}
