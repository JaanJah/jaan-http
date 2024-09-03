package paths

import "github.com/gofiber/fiber/v3"

func OK(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}