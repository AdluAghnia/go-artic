package handler

import (
	"log"
	"time"

	"github.com/AdluAghnia/go-artic/middleware"
	"github.com/AdluAghnia/go-artic/models"
	"github.com/gofiber/fiber/v2"
)

func LoginViewHandler(c *fiber.Ctx) error {
	return c.Render("login", nil, "layouts/main")
}

func LoginHandler(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := models.GetUserByEmail(email)
	if err != nil {
		log.Println(err)
		return c.SendString(err.Error())
	}

	ok, err := user.ComparePassword(password)
	if err != nil {
		return c.SendString(err.Error())
	}

	if !ok {
		return c.SendString("Please check your password or email")
	}


	// Generate JWT Token
	token, err := middleware.GenerateJWTKey(user)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Couldn't generate token")
	}

	// Create Cookie
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour),
	})

	if c.Get("HX-Request") == "true" {
		c.Set("HX-Redirect", "/")
		return nil
	}

	return c.Redirect("/", fiber.StatusSeeOther)
}

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
