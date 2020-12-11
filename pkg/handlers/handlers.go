package handlers

import "net/http"

func GetData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("data"))
	}
}
