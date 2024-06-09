package main

import (
	"fmt"
	"net/http"

	"github.com/coderharsx1122/rssscr/auth"
	"github.com/coderharsx1122/rssscr/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cnf *apiConfig) authMiddleware(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("error:%v", err))
		}

		user, err := cnf.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("error:%v", err))
		}

		handler(w, r, user)
	}
}
