package middleware

import (
	"Weather-API-GoExpertPostGrad/pkg/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func CheckCepMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cep := chi.URLParam(r, "cep")
		if cep == "" || len(cep) == 0 {
			http.Error(w, "CEP is required", http.StatusBadRequest)
			return
		}

		if !utils.IsValidZipcode(cep) {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}
		next.ServeHTTP(w, r)
	})
}
