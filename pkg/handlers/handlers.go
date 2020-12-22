package handlers

import (
	"net/http"

	mw "github.com/nikimanoledaki/tracing-tutorial/pkg/tracing"
	"github.com/opentracing/opentracing-go/log"
)

// GetData is an example handler func.
func GetData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		span, begin := mw.StartChildSpan(r.Context(), "get_data")
		defer mw.FinishChildSpan(span, begin)

		// The span ends once this function returns, with or
		// without errors. It logs any db-related errors
		// before returning.

		// Setting tags here
		span.SetTag("kind", "client")
		span.SetTag("db.type", "presto")
		span.SetTag("db.statement", "SELECT * FROM db")

		rows, err := s.Db.Query(stmt, args...)
		if err != nil {
			log.Error(err)
		}
	}
}
