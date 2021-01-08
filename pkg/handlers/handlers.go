package handlers

import (
	"net/http"

	mw "github.com/nikimanoledaki/tracing-tutorial/pkg/tracing"
)

// GetData is an example handler func.
func GetData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		span, begin := mw.StartChildSpan(r.Context(), "get_data")
		defer mw.FinishChildSpan(span, begin)

		// The span ends once this function returns, with or
		// without errors. It logs any db-related errors
		// before returning.

		// Set tags here
		span.SetTag("kind", "client")
		span.SetTag("path", r.URL.Path)
	}
}
