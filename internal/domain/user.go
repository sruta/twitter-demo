package domain

import (
	"time"
)

type User struct {
	ID        int64
	Email     string
	Password  string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
