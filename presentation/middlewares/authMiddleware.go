package middlewares

import (
	"fmt"
	"io"
	"main/usecases"
	"net/http"
	"strings"
)

func Auth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	var t string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjE3MDYzNzQ4NTQsInN1YiI6ImUzYWUwZGJkLTU3NDUtNGY4Ny1hN2E1LWU4ZjJhZDU3NjMzZiJ9.sB5i9GH5jJY6KsYBDLEKe5rClAdPrVqUP7Gt9bXu3rw"

	return func(w http.ResponseWriter, r *http.Request) {
		verify, err := usecases.ValidateToken(t)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
			return
		}

		id, err := verify.GetSubject()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
			return
		}

		newBody := fmt.Sprintf(`{"id": %s, "body": %s}`, id, string(b))

		r.Body = io.NopCloser(strings.NewReader(newBody))
		defer r.Body.Close()
		fmt.Println("middleware", id)
		handlerFunc.ServeHTTP(w, r)
	}
}

func AuthTwo(handler http.Handler) http.Handler {
	var t string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjE3MDY4MTc0NzMsInN1YiI6IjNiY2I3YzZhLWRmMDctNGNlZi1hMGMyLWUzNjEyZDBiZDY2ZiJ9.tCwpaM9gshz3Wv0hcWHDkIfEv4xCssECrJGPjaV54aU"

	fn := func(w http.ResponseWriter, r *http.Request) {

		verify, err := usecases.ValidateToken(t)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
			return
		}

		id, err := verify.GetSubject()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
			return
		}
		fmt.Println("middleware", id)
		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
