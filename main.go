package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikimanoledaki/tracing-tutorial/cmd"
	"github.com/nikimanoledaki/tracing-tutorial/pkg/tracing"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	closer, err := cmd.InitJaeger()
	if err != nil {
		log.Error(err)
	}
	defer closer.Close()

	r := mux.NewRouter()
	r.Use(tracing.ParentSpan)
	r.Handle("/data", getData())

	_ = http.ListenAndServe(":9092", r)
}

func getData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("data"))
	}
}
