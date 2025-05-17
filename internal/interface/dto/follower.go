package dto

import (
	"time"
	"twitter-uala/internal/domain"
)

type FollowerCreate struct {
	FollowerID int64 `json:"follower_id" validate:"required"`
	FollowedID int64 `json:"followed_id" validate:"required"`
}

type FollowerResponse struct {
	FollowerID int64     `json:"follower_id"`
	FollowedID int64     `json:"followed_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func FromFollowerCreateToFollower(dto FollowerCreate) domain.Follower {
	return domain.Follower{
		FollowerID: dto.FollowerID,
		FollowedID: dto.FollowedID,
	}
}

func FromFollowerToFollowerResponse(follower domain.Follower) FollowerResponse {
	return FollowerResponse{
		FollowerID: follower.FollowerID,
		FollowedID: follower.FollowedID,
		CreatedAt:  follower.CreatedAt,
	}
}
