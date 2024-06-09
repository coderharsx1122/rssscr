package main

import (
	"time"

	"github.com/coderharsx1122/rssscr/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}
type FollowFeed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

func dbFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func dbFeedsToFeeds(dfeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, feed := range dfeeds {
		feeds = append(feeds, dbFeedToFeed(feed))
	}
	return feeds
}
func dbFollowFeedToFollowFeed(dfeed database.FeedFollow) FollowFeed {
	return FollowFeed{
		ID:        dfeed.ID,
		CreatedAt: dfeed.CreatedAt,
		UpdatedAt: dfeed.UpdatedAt,
		UserID:    dfeed.UserID,
		FeedID:    dfeed.FeedID,
	}
}
func dbFeedFollowsToFeedFollows(dfeeds []database.FeedFollow) []FollowFeed {
	feeds := []FollowFeed{}

	for _, feedFollow := range dfeeds {
		feeds = append(feeds, dbFollowFeedToFollowFeed(feedFollow))
	}

	return feeds
}
