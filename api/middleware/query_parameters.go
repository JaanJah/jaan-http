package middleware

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

func QueryParameters(c fiber.Ctx) error {
	timeoutParam := c.Query("timeout")

	if timeoutParam != "" {
		intValue, err := strconv.Atoi(timeoutParam)
		if err != nil {
			log.Printf("Error converting timeout parameter to integer: %s", err)
			return c.Next()
		}
		time.Sleep(time.Duration(intValue) * time.Millisecond)
	}

	return c.Next()
}
