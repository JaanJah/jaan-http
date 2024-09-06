package middleware

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

func QueryParameters(c fiber.Ctx) error {
	delayParam := c.Query("delay")

	if delayParam != "" {
		intValue, err := strconv.Atoi(delayParam)
		if err != nil {
			log.Printf("Error converting delay parameter to integer: %s", err)
			return c.Next()
		}
		time.Sleep(time.Duration(intValue) * time.Millisecond)
	}

	return c.Next()
}
