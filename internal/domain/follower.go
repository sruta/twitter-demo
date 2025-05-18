package domain

import "time"

type Follower struct {
	FollowerID int64
	FollowedID int64
	CreatedAt  time.Time
}
