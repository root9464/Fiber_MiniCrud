package main

import (
	"log"
	database "root/src/Database"
	routes "root/src/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func AllRoutes(app *fiber.App) {
	app.Get("/", routes.Hello)
	app.Post("/adduser", routes.AddUser)
	app.Get("/allusers", routes.AllUsers)
}

func main() {
	database.ConnectDb()
    app := fiber.New()
	AllRoutes(app)
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}