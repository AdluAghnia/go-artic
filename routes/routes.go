package routes

import (
	"github.com/AdluAghnia/go-artic/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/", handler.IndexHandler)
	app.Get("/article/create", handler.CreateArticleHandler)
	app.Get("/article/:id", handler.ShowArticle)
	app.Post("/article", handler.SaveArticleHandler)

	app.Get("/register", handler.RegisterViewHandler)
	app.Post("/auth/register", handler.RegisterHandler)
}
