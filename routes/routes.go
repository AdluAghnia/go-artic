package routes

import (
	"github.com/AdluAghnia/go-artic/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/", handler.IndexHandler)
	app.Get("/article/create", handler.CreateArticleHandler)

	app.Post("/article", handler.SaveArticleHandler)
}
