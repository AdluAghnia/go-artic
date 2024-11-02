package routes

import (
	"github.com/AdluAghnia/go-artic/handler"
	"github.com/AdluAghnia/go-artic/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/", middleware.JWTMiddleware(), handler.IndexHandler)
	app.Get("/article/create", middleware.JWTMiddleware(), handler.CreateArticleHandler)
	app.Get("/article/:id", middleware.JWTMiddleware(), handler.ShowArticle)
	app.Post("/article", middleware.JWTMiddleware(), handler.SaveArticleHandler)

	app.Get("/register", handler.RegisterViewHandler)
	app.Post("/auth/register", handler.RegisterHandler)
}
