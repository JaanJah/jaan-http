package main

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestSetupStatusCodeRoutes(t *testing.T) {
	app := fiber.New()
	SetupStatusCodeRoutes(app)

	tests := []struct {
		name       string
		method     string
		path       string
		statusCode int
	}{
		{"GET /100", "GET", "/100", fiber.StatusContinue},
		{"POST /200", "POST", "/200", fiber.StatusOK},
		{"PUT /201", "PUT", "/201", fiber.StatusCreated},
		{"DELETE /202", "DELETE", "/202", fiber.StatusAccepted},
		{"GET /404", "GET", "/404", fiber.StatusNotFound},
		{"GET /500", "GET", "/500", fiber.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			res, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}

			if res.StatusCode != tt.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					res.StatusCode, tt.statusCode)
			}
		})
	}
}
