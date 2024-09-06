package middleware

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
)

func QueryParameters(c fiber.Ctx) error {
	delayQuery(c.Query("delay"))

	return c.Next()
}

func delayQuery(delay string) {
	if delay == "" {
		return
	}

	intValue, err := strconv.Atoi(delay)
	if err != nil {
		// TODO: Figure out proper way how to handle logs and errors like this
		log.Printf("Error converting delay parameter to integer: %s", err)
		return
	}
	log.Printf("Delaying response by %d milliseconds", intValue)
	time.Sleep(time.Duration(intValue) * time.Millisecond)
}
