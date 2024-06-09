package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/coderharsx1122/rssscr/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (cnf *apiConfig) feedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Feed_id uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error: %v", err))
	}

	followfeed, err := cnf.DB.CreateFollowFeed(r.Context(), database.CreateFollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.Feed_id,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't create follow feed: %v", err))
		return
	}

	respond(w, 200, dbFollowFeedToFollowFeed(followfeed))
}

func (cnf *apiConfig) getFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	followfeeds, err := cnf.DB.GetFollowFeed(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't create follow feed: %v", err))
		return
	}
	respond(w, 200, dbFeedFollowsToFeedFollows(followfeeds))
}

func (cnf *apiConfig) deleteFeedFollowHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowIdstr := chi.URLParam(r, "followFeedId")

	feedFollowId, err := uuid.Parse(feedFollowIdstr)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't delete follow feed: %v", err))
		return
	}

	errr := cnf.DB.DeleteFollowFeed(r.Context(), database.DeleteFollowFeedParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})

	if errr != nil {
		respondWithError(w, 400, fmt.Sprintf("Could't delete follow feed: %v", err))
		return
	}

	type message struct {
		message string
	}
	respond(w, 200, struct{}{})
}
