package servlet

type (
	// RequestMiddleware defines a function to process middleware.
	RequestMiddleware = func(RequestHandler) RequestHandler

	// RequestHandler defines a function to serve HTTP requests.
	RequestHandler = func(RequestContext) error

	// HTTPErrorHandler is a centralized HTTP error handler.
	HTTPErrorHandler = func(error, RequestContext)
)
