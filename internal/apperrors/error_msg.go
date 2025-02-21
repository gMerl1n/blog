package apperrors

var (
	// 4xx client error
	IncorrectRequest     = newError("incorrect request", 400)
	IncorrectRequestBody = newError("incorrect request body", 400)
	IncorrectData        = newError("data retrieval error", 400)

	RecordingError = newError("recording error", 400)
	NoPostFound    = newError("no post found", 404)

	IncorrectRequestParams = newError("incorrect request parametrs", 400)
	Unauthorized           = newError("unauthorized", 401)
	Forbidden              = newError("forbidden", 403)
	NotFound               = newError("not found", 404)
	NotAllowed             = newError("method not allowed", 405)
	Conflict               = newError("conflict", 409)
	PostIsAlready          = newError("this post already exists", 409)
	Gone                   = newError("requested resource is no longer available", 410)
	PayloadTooLarge        = newError("payload too large", 413)
	TooManyRequests        = newError("too many requests", 429)
	ClientClosedRequest    = newError("client closed request", 499)

	// 5xx server error
	InternalServerError = newError("internal server error", 500)
	NotImplemented      = newError("method is not implemented", 501)
	BadGateway          = newError("server received invalid response from upstream", 502)
	ServiceUnavailable  = newError("service is not available", 503)
	UnknownError        = newError("unknown error", 520)
)
