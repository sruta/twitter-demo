package domain

import "time"

type Tweet struct {
	ID        int64
	UserID    int64
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User
}
