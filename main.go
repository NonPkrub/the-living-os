package main

import (
	"short-url/adapters"
	"short-url/config"
	"short-url/entities"
	"short-url/usecases"

	_ "short-url/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title ShortURL API
// @version 1.0
// @description This is an API server for short url.
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	db, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	urlRepo := adapters.NewGormShortUrlRepository(db)
	urlUseCase := usecases.NewShortURLService(urlRepo)
	httpShortUrlHandler := adapters.NewHttpShortUrlHandler(urlUseCase)

	api := app.Group("/api")
	api.Post("/shorten", httpShortUrlHandler.CreateShortUrl)
	api.Get(":shortCode", httpShortUrlHandler.Redirect)
	api.Get("/stats/:shortCode", httpShortUrlHandler.GetShortUrl)

	app.Get("/docs/*", swagger.HandlerDefault)

	db.AutoMigrate(&entities.URLShorten{})
	app.Listen(":3000")
}
