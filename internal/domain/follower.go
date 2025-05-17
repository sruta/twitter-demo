package domain

import "time"

type Follower struct {
	FollowerID int64     `json:"follower_id"`
	FollowedID int64     `json:"followed_id"`
	CreatedAt  time.Time `json:"created_at"`
}
