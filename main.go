package main

import (
	"fmt"
	"log"

	"github.com/AdluAghnia/go-artic/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./templates/", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.SetupRouter(app)
	fmt.Println("I'm Batman")

	log.Fatal(app.Listen(":8080"))
}
