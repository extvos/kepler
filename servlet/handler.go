package servlet

type (
	// MiddlewareFunc defines a function to process middleware.
	MiddlewareFunc func(HandlerFunc) HandlerFunc

	// HandlerFunc defines a function to serve HTTP requests.
	HandlerFunc func(RequestContext) error

	// HTTPErrorHandler is a centralized HTTP error handler.
	HTTPErrorHandler func(error, RequestContext)
)
