package handler

import (
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.Render("index", nil, "layouts/main")
}

func CreateArticleHandler(c *fiber.Ctx) error {
	return c.Render("create-article", nil)
}
