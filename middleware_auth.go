package main

import (
	"fmt"
	"net/http"

	"github.com/Waffenlord/blog-aggregator/internal/auth"
	"github.com/Waffenlord/blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)


func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 400, fmt.Sprintf("couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}