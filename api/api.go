package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JaanJah/jaan-http/api/middleware"
	"github.com/JaanJah/jaan-http/api/status_code"
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

	app.Use(middleware.QueryParameters)

	SetupStatusCodeRoutes(app)

	app.Listen(listenerAddress, fiber.ListenConfig{
		EnablePrintRoutes: os.Getenv("ENABLE_PRINT_ROUTES") == "true",
	})
}

func SetupStatusCodeRoutes(app *fiber.App) {
	app.Use("/100", status_code.StatusContinue)
	app.Use("/101", status_code.StatusSwitchingProtocols)
	app.Use("/102", status_code.StatusProcessing)
	app.Use("/103", status_code.StatusEarlyHints)
	app.Use("/200", status_code.StatusOK)
	app.Use("/201", status_code.StatusCreated)
	app.Use("/202", status_code.StatusAccepted)
	app.Use("/203", status_code.StatusNonAuthoritativeInformation)
	app.Use("/204", status_code.StatusNoContent)
	app.Use("/205", status_code.StatusResetContent)
	app.Use("/206", status_code.StatusPartialContent)
	app.Use("/207", status_code.StatusMultiStatus)
	app.Use("/208", status_code.StatusAlreadyReported)
	app.Use("/226", status_code.StatusIMUsed)
	app.Use("/300", status_code.StatusMultipleChoices)
	app.Use("/301", status_code.StatusMovedPermanently)
	app.Use("/302", status_code.StatusFound)
	app.Use("/303", status_code.StatusSeeOther)
	app.Use("/304", status_code.StatusNotModified)
	app.Use("/305", status_code.StatusUseProxy)
	app.Use("/306", status_code.StatusSwitchProxy)
	app.Use("/307", status_code.StatusTemporaryRedirect)
	app.Use("/308", status_code.StatusPermanentRedirect)
	app.Use("/400", status_code.StatusBadRequest)
	app.Use("/401", status_code.StatusUnauthorized)
	app.Use("/402", status_code.StatusPaymentRequired)
	app.Use("/403", status_code.StatusForbidden)
	app.Use("/404", status_code.StatusNotFound)
	app.Use("/405", status_code.StatusMethodNotAllowed)
	app.Use("/406", status_code.StatusNotAcceptable)
	app.Use("/407", status_code.StatusProxyAuthenticationRequired)
	app.Use("/408", status_code.StatusRequestTimeout)
	app.Use("/409", status_code.StatusConflict)
	app.Use("/410", status_code.StatusGone)
	app.Use("/411", status_code.StatusLengthRequired)
	app.Use("/412", status_code.StatusPreconditionFailed)
	app.Use("/413", status_code.StatusPayloadTooLarge)
	app.Use("/414", status_code.StatusURITooLong)
	app.Use("/415", status_code.StatusUnsupportedMediaType)
	app.Use("/416", status_code.StatusRangeNotSatisfiable)
	app.Use("/417", status_code.StatusExpectationFailed)
	app.Use("/418", status_code.StatusImATeapot)
	app.Use("/421", status_code.StatusMisdirectedRequest)
	app.Use("/422", status_code.StatusUnprocessableEntity)
	app.Use("/423", status_code.StatusLocked)
	app.Use("/424", status_code.StatusFailedDependency)
	app.Use("/425", status_code.StatusTooEarly)
	app.Use("/426", status_code.StatusUpgradeRequired)
	app.Use("/428", status_code.StatusPreconditionRequired)
	app.Use("/429", status_code.StatusTooManyRequests)
	app.Use("/431", status_code.StatusRequestHeaderFieldsTooLarge)
	app.Use("/451", status_code.StatusUnavailableForLegalReasons)
	app.Use("/500", status_code.StatusInternalServerError)
	app.Use("/501", status_code.StatusNotImplemented)
	app.Use("/502", status_code.StatusBadGateway)
	app.Use("/503", status_code.StatusServiceUnavailable)
	app.Use("/504", status_code.StatusGatewayTimeout)
	app.Use("/505", status_code.StatusHTTPVersionNotSupported)
	app.Use("/506", status_code.StatusVariantAlsoNegotiates)
	app.Use("/507", status_code.StatusInsufficientStorage)
	app.Use("/508", status_code.StatusLoopDetected)
	app.Use("/510", status_code.StatusNotExtended)
	app.Use("/511", status_code.StatusNetworkAuthenticationRequired)
}
