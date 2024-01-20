package middlewares

import (
	"fmt"
	"main/usecases"
	"net/http"
)

func Auth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	var t string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjE3MDU3MDkxMzAsInN1YiI6ImUzYWUwZGJkLTU3NDUtNGY4Ny1hN2E1LWU4ZjJhZDU3NjMzZiJ9.pYO_QvFyMygviQSH9nqC8liv8sFVos15IHnBetd-y1M"

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

		fmt.Println("middleware", id)
		handlerFunc.ServeHTTP(w, r)
	}
}

func AuthTwo(handler http.Handler) http.Handler {
	var t string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjE3MDU3MDg0MzEsInN1YiI6ImUzYWUwZGJkLTU3NDUtNGY4Ny1hN2E1LWU4ZjJhZDU3NjMzZiJ9.w7cmGZclVdBZSNgf7btC5IQSNg4dMX8uIIgZIskQINc"

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
		return
	}

	return http.HandlerFunc(fn)
}
