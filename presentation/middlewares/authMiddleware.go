package middlewares

import (
	"fmt"
	"main/usecases"
	"net/http"
)

func Auth(handlerFunc http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		token := r.Header.Get("Authorization")
		_, err := usecases.ValidateToken(token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
			return
		}

		handlerFunc.ServeHTTP(w, r)
	}
}

func GetUserId(w http.ResponseWriter, r *http.Request) string {
	token := r.Header.Get("Authorization")

	verify, _ := usecases.ValidateToken(token)
	id, _ := verify.GetSubject()

	return id
}
