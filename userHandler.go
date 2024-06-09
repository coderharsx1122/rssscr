package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/coderharsx1122/rssscr/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error: %v", err))
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't create user: %v", err))
		return
	}

	respond(w, 200, dbUserToUser(user))
}

// get user handler
func (apiCfg *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	respond(w, 200, dbUserToUser(user))
}
