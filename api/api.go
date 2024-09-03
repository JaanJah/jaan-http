package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JaanJah/jaan-http/api/paths"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	app := fiber.New()
	listenerAddress := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	app.Get("/", paths.OK)

	app.Listen(listenerAddress, fiber.ListenConfig{
		EnablePrintRoutes: true,
	})
}
