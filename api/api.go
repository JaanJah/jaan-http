package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JaanJah/jaan-http/api/status_codes"
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

	// Status code paths
	app.Use("/", status_codes.StatusOK)
	app.Use("/100", status_codes.StatusContinue)
	app.Use("/101", status_codes.StatusSwitchingProtocols)
	app.Use("/102", status_codes.StatusProcessing)
	app.Use("/103", status_codes.StatusEarlyHints)
	app.Use("/200", status_codes.StatusOK)
	app.Use("/201", status_codes.StatusCreated)
	app.Use("/202", status_codes.StatusAccepted)
	app.Use("/203", status_codes.StatusNonAuthoritativeInformation)
	app.Use("/204", status_codes.StatusNoContent)
	app.Use("/205", status_codes.StatusResetContent)
	app.Use("/206", status_codes.StatusPartialContent)
	app.Use("/207", status_codes.StatusMultiStatus)
	app.Use("/208", status_codes.StatusAlreadyReported)
	app.Use("/226", status_codes.StatusIMUsed)
	app.Use("/300", status_codes.StatusMultipleChoices)
	app.Use("/301", status_codes.StatusMovedPermanently)
	app.Use("/302", status_codes.StatusFound)
	app.Use("/303", status_codes.StatusSeeOther)
	app.Use("/304", status_codes.StatusNotModified)
	app.Use("/305", status_codes.StatusUseProxy)
	app.Use("/306", status_codes.StatusSwitchProxy)
	app.Use("/307", status_codes.StatusTemporaryRedirect)
	app.Use("/308", status_codes.StatusPermanentRedirect)
	app.Use("/400", status_codes.StatusBadRequest)
	app.Use("/401", status_codes.StatusUnauthorized)
	app.Use("/402", status_codes.StatusPaymentRequired)
	app.Use("/403", status_codes.StatusForbidden)
	app.Use("/404", status_codes.StatusNotFound)
	app.Use("/405", status_codes.StatusMethodNotAllowed)
	app.Use("/406", status_codes.StatusNotAcceptable)
	app.Use("/407", status_codes.StatusProxyAuthenticationRequired)
	app.Use("/408", status_codes.StatusRequestTimeout)
	app.Use("/409", status_codes.StatusConflict)
	app.Use("/410", status_codes.StatusGone)
	app.Use("/411", status_codes.StatusLengthRequired)
	app.Use("/412", status_codes.StatusPreconditionFailed)
	app.Use("/413", status_codes.StatusPayloadTooLarge)
	app.Use("/414", status_codes.StatusURITooLong)
	app.Use("/415", status_codes.StatusUnsupportedMediaType)
	app.Use("/416", status_codes.StatusRangeNotSatisfiable)
	app.Use("/417", status_codes.StatusExpectationFailed)
	app.Use("/418", status_codes.StatusImATeapot)
	app.Use("/421", status_codes.StatusMisdirectedRequest)
	app.Use("/422", status_codes.StatusUnprocessableEntity)
	app.Use("/423", status_codes.StatusLocked)
	app.Use("/424", status_codes.StatusFailedDependency)
	app.Use("/425", status_codes.StatusTooEarly)
	app.Use("/426", status_codes.StatusUpgradeRequired)
	app.Use("/428", status_codes.StatusPreconditionRequired)
	app.Use("/429", status_codes.StatusTooManyRequests)
	app.Use("/431", status_codes.StatusRequestHeaderFieldsTooLarge)
	app.Use("/451", status_codes.StatusUnavailableForLegalReasons)
	app.Use("/500", status_codes.StatusInternalServerError)
	app.Use("/501", status_codes.StatusNotImplemented)
	app.Use("/502", status_codes.StatusBadGateway)
	app.Use("/503", status_codes.StatusServiceUnavailable)
	app.Use("/504", status_codes.StatusGatewayTimeout)
	app.Use("/505", status_codes.StatusHTTPVersionNotSupported)
	app.Use("/506", status_codes.StatusVariantAlsoNegotiates)
	app.Use("/507", status_codes.StatusInsufficientStorage)
	app.Use("/508", status_codes.StatusLoopDetected)
	app.Use("/510", status_codes.StatusNotExtended)
	app.Use("/511", status_codes.StatusNetworkAuthenticationRequired)

	app.Listen(listenerAddress, fiber.ListenConfig{
		EnablePrintRoutes: os.Getenv("ENABLE_PRINT_ROUTES") == "true",
	})
}
