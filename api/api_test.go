package main

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v3"
)


func TestSetupStatusCodeRoutes(t *testing.T) {
	app := fiber.New()
	SetupStatusCodeRoutes(app)

	req, err := http.NewRequest("GET", "/500", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code.
	if res.StatusCode != fiber.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			res.StatusCode, fiber.StatusInternalServerError)
	}
}
