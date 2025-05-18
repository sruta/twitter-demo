package dto

import (
	"time"
	"twitter-uala/internal/domain"
)

type TweetCreate struct {
	UserID int64  `json:"user_id" validate:"required"`
	Text   string `json:"text" validate:"required,min=1,max=280"`
}

type TweetUpdate struct {
	ID   int64  `json:"id" validate:"required"`
	Text string `json:"text" validate:"required,min=1,max=280"`
}

type TweetResponse struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromTweetCreateToTweet(dto TweetCreate) domain.Tweet {
	return domain.Tweet{
		UserID: dto.UserID,
		Text:   dto.Text,
	}
}

func FromTweetUpdateToTweet(dto TweetUpdate) domain.Tweet {
	return domain.Tweet{
		ID:   dto.ID,
		Text: dto.Text,
	}
}

func FromTweetToTweetResponse(tweet domain.Tweet) TweetResponse {
	return TweetResponse{
		ID:        tweet.ID,
		UserID:    tweet.UserID,
		Text:      tweet.Text,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
	}
}
