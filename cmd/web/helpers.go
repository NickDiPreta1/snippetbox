package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError logs detailed error information (including a stack trace)
// and sends a generic "500 Internal Server Error" response to the client.
//
// This should be used for unexpected server-side errors (e.g. bugs, file I/O failures).
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends a specific client-side error response (4xx) to the user.
//
// This should be used when the client makes a bad request — such as malformed input,
// missing parameters, or unauthorized access. It sends the standard HTTP status text
// (e.g., "Bad Request", "Unauthorized", "Not Found") as the response body.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound is a convenience wrapper that sends a "404 Not Found" response.
//
// It’s just a shorthand for app.clientError(w, http.StatusNotFound)
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
