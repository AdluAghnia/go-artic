package handler

import (
	"github.com/AdluAghnia/go-artic/models"
	"github.com/gofiber/fiber/v2"
)

func RegisterViewHandler(c *fiber.Ctx) error {
	return c.Render("register", nil, "layouts/main")
}

func RegisterHandler(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirm_password := c.FormValue("confirm_password")

	if confirm_password != password {
		return c.SendString("Salah password")
	}

	user := models.NewUser(
		username,
		email,
		password,
	)

	isValid, errs := user.ValidateRegisterUser()
	if isValid == false {
		return c.Render("register", fiber.Map{
			"Errors": errs,
		})
	}

	err := user.SaveUser()
	if err != nil {
		return c.SendString(err.Error())
	}
	
	if c.Get("HX-Request") == "true" {
		c.Set("HX-Redirect", "/login")
		return nil
	}

	return c.Redirect("/login", fiber.StatusSeeOther)
}
