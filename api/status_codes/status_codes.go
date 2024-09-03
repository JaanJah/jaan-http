package status_codes

import "github.com/gofiber/fiber/v3"

// 100 Continue
func StatusContinue(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusContinue)
}

// 101 Switching Protocols
func StatusSwitchingProtocols(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusSwitchingProtocols)
}

// 102 Processing
func StatusProcessing(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusProcessing)
}

// 103 Early Hints
func StatusEarlyHints(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusEarlyHints)
}

// 200 OK
func StatusOK(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// 201 Created
func StatusCreated(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusCreated)
}

// 202 Accepted
func StatusAccepted(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusAccepted)
}

// 203 Non-Authoritative Information
func StatusNonAuthoritativeInformation(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNonAuthoritativeInformation)
}

// 204 No Content
func StatusNoContent(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

// 205 Reset Content
func StatusResetContent(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusResetContent)
}

// 206 Partial Content
func StatusPartialContent(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusPartialContent)
}

// 207 Multi-Status
func StatusMultiStatus(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusMultiStatus)
}

// 208 Already Reported
func StatusAlreadyReported(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusAlreadyReported)
}

// 226 IM Used
func StatusIMUsed(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusIMUsed)
}

// 300 Multiple Choices
func StatusMultipleChoices(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusMultipleChoices)
}

// 301 Moved Permanently
func StatusMovedPermanently(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusMovedPermanently)
}

// 302 Found
func StatusFound(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusFound)
}

// 303 See Other
func StatusSeeOther(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusSeeOther)
}

// 304 Not Modified
func StatusNotModified(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotModified)
}

// 305 Use Proxy
func StatusUseProxy(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusUseProxy)
}

// 306 Switch Proxy
func StatusSwitchProxy(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusSwitchProxy)
}

// 307 Temporary Redirect
func StatusTemporaryRedirect(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusTemporaryRedirect)
}

// 308 Permanent Redirect
func StatusPermanentRedirect(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusPermanentRedirect)
}

// 400 Bad Request
func StatusBadRequest(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusBadRequest)
}

// 401 Unauthorized
func StatusUnauthorized(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusUnauthorized)
}

// 402 Payment Required
func StatusPaymentRequired(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusPaymentRequired)
}

// 403 Forbidden
func StatusForbidden(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusForbidden)
}

// 404 Not Found
func StatusNotFound(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotFound)
}

// 405 Method Not Allowed
func StatusMethodNotAllowed(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusMethodNotAllowed)
}

// 406 Not Acceptable
func StatusNotAcceptable(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotAcceptable)
}

// 407 Proxy Authentication Required
func StatusProxyAuthenticationRequired(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusProxyAuthRequired)
}

// 408 Request Timeout
func StatusRequestTimeout(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusRequestTimeout)
}

// 409 Conflict
func StatusConflict(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusConflict)
}

// 410 Gone
func StatusGone(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusGone)
}

// 411 Length Required
func StatusLengthRequired(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusLengthRequired)
}

// 412 Precondition Failed
func StatusPreconditionFailed(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusPreconditionFailed)
}

// 413 Payload Too Large
func StatusPayloadTooLarge(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusRequestEntityTooLarge)
}

// 414 URI Too Long
func StatusURITooLong(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusRequestURITooLong)
}

// 415 Unsupported Media Type
func StatusUnsupportedMediaType(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusUnsupportedMediaType)
}

// 416 Range Not Satisfiable
func StatusRangeNotSatisfiable(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusRequestedRangeNotSatisfiable)
}

// 417 Expectation Failed
func StatusExpectationFailed(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusExpectationFailed)
}

// 418 I'm a teapot
func StatusImATeapot(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusTeapot)
}

// 421 Misdirected Request
func StatusMisdirectedRequest(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusMisdirectedRequest)
}

// 422 Unprocessable Entity
func StatusUnprocessableEntity(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusUnprocessableEntity)
}

// 423 Locked
func StatusLocked(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusLocked)
}

// 424 Failed Dependency
func StatusFailedDependency(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusFailedDependency)
}

// 425 Too Early
func StatusTooEarly(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusTooEarly)
}

// 426 Upgrade Required
func StatusUpgradeRequired(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusUpgradeRequired)
}

// 428 Precondition Required
func StatusPreconditionRequired(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusPreconditionRequired)
}

// 429 Too Many Requests
func StatusTooManyRequests(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusTooManyRequests)
}

// 431 Request Header Fields Too Large
func StatusRequestHeaderFieldsTooLarge(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusRequestHeaderFieldsTooLarge)
}

// 451 Unavailable For Legal Reasons
func StatusUnavailableForLegalReasons(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusUnavailableForLegalReasons)
}

// 500 Internal Server Error
func StatusInternalServerError(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusInternalServerError)
}

// 501 Not Implemented
func StatusNotImplemented(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

// 502 Bad Gateway
func StatusBadGateway(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusBadGateway)
}

// 503 Service Unavailable
func StatusServiceUnavailable(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusServiceUnavailable)
}

// 504 Gateway Timeout
func StatusGatewayTimeout(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusGatewayTimeout)
}

// 505 HTTP Version Not Supported
func StatusHTTPVersionNotSupported(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusHTTPVersionNotSupported)
}

// 506 Variant Also Negotiates
func StatusVariantAlsoNegotiates(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusVariantAlsoNegotiates)
}

// 507 Insufficient Storage
func StatusInsufficientStorage(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusInsufficientStorage)
}

// 508 Loop Detected
func StatusLoopDetected(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusLoopDetected)
}

// 510 Not Extended
func StatusNotExtended(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotExtended)
}

// 511 Network Authentication Required
func StatusNetworkAuthenticationRequired(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNetworkAuthenticationRequired)
}