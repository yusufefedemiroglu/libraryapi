package main

import (
	"fmt"
	"go-library/database"
	"go-library/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	client, err := database.Connect()
	fmt.Println(err)
	defer database.Disconnect()

	database.InitCollections(client)

	app := fiber.New()
	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Authorization, Content-Type, x-user-id",
		AllowCredentials: true,
	}))
	app.Post("/users", handlers.CreateUser)
	app.Post("/books", handlers.CreateBook)
	app.Get("/books", handlers.GetBooks)
	app.Post("/login", handlers.HandleLogin)
	app.Post("/create", handlers.CreateUser)

	app.Listen(":3000")
}
