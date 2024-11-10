package main

import (
	"log"
	"os"

	"example.com/go-fiber-crud/db"
	"example.com/go-fiber-crud/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	db.InitDB()

	port := os.Getenv("PORT")

	app := fiber.New()

	routes.RegisterRoutes(app)

	app.Listen(":" + port)
}
