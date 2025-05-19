package dto

import (
	"time"
	"twitter-demo/internal/domain"
)

type TweetCreate struct {
	UserID int64  `json:"user_id" validate:"required"`
	Text   string `json:"text" validate:"required,min=1,max=280"`
}

type TweetUpdate struct {
	ID     int64  `json:"id" validate:"required"`
	UserID int64  `json:"user_id" validate:"required"`
	Text   string `json:"text" validate:"required,min=1,max=280"`
}

type TweetResponse struct {
	ID        int64         `json:"id"`
	UserID    int64         `json:"user_id"`
	Text      string        `json:"text"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      *UserResponse `json:"user,omitempty"`
}

func FromTweetCreateToTweet(dto TweetCreate) domain.Tweet {
	return domain.Tweet{
		UserID: dto.UserID,
		Text:   dto.Text,
	}
}

func FromTweetUpdateToTweet(dto TweetUpdate) domain.Tweet {
	return domain.Tweet{
		ID:     dto.ID,
		UserID: dto.UserID,
		Text:   dto.Text,
	}
}

func FromTweetToTweetResponse(tweet domain.Tweet) TweetResponse {
	var user *UserResponse
	if tweet.User != nil {
		userResponse := FromUserToUserResponse(*tweet.User)
		user = &userResponse
	}

	return TweetResponse{
		ID:        tweet.ID,
		UserID:    tweet.UserID,
		Text:      tweet.Text,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: tweet.UpdatedAt,
		User:      user,
	}
}

func FromTweetsToTweetsResponse(tweets []domain.Tweet) []TweetResponse {
	response := []TweetResponse{}
	for _, tweet := range tweets {
		response = append(response, FromTweetToTweetResponse(tweet))
	}
	return response
}
