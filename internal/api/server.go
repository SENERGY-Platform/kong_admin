package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Handler is an interface providing a generic mechanism for handling HTTP requests.
type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request) error
}

// HTTPHandler is a wrapper that encapsulates Handler interface as http.Handler.
// HTTPHandler also enforces that the Handler only responds to requests with registered HTTP methods.
type HTTPHandler struct {
	Handler          // CFSSL handler
	Methods []string // The associated HTTP methods
}

// HandlerFunc is similar to the http.HandlerFunc type; it serves as
// an adapter allowing the use of ordinary functions as Handlers. If
// f is a function with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.
type HandlerFunc func(http.ResponseWriter, *http.Request) error

// Handle calls f(w, r)
func (f HandlerFunc) Handle(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	return f(w, r)
}

// HandleError is the centralised error handling and reporting.
func HandleError(w http.ResponseWriter, err error) (code int) {
	if err == nil {
		return http.StatusOK
	}
	msg := err.Error()
	httpCode := http.StatusInternalServerError

	// If it is recognized as HttpError emitted from cfssl,
	// we rewrite the status code accordingly. If it is a
	// cfssl error, set the http status to StatusBadRequest
	switch err := err.(type) {
	case *HTTPError:
		httpCode = err.StatusCode
		code = err.StatusCode
	}

	jsonMessage, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
	} else {
		msg = string(jsonMessage)
	}
	http.Error(w, msg, httpCode)
	return code
}

// ServeHTTP encapsulates the call to underlying Handler to handle the request
// and return the response with proper HTTP status code
func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var match bool
	// Throw 405 when requested with an unsupported verb.
	for _, m := range h.Methods {
		if m == r.Method {
			match = true
		}
	}
	if match {
		err = h.Handle(w, r)
	} else {
		err = NewMethodNotAllowed(r.Method)
	}
	status := HandleError(w, err)
	log.Printf("%s - \"%s %s\" %d", r.RemoteAddr, r.Method, r.URL, status)
}
