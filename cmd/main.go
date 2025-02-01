package main

import (
	"log"

	"github.com/MarcKVR/clean_arquitecture/cmd/config"
	"github.com/MarcKVR/clean_arquitecture/infrastructure/database"
	"github.com/MarcKVR/clean_arquitecture/presentation/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading config initializatin: %v", err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡CONGRATULATIONS! Welcome to te API services!")
	})

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	defer database.Close(db)

	routes.RegisterRoutes(app, db, cfg)

	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
