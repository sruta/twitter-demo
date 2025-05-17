package domain

import (
	"encoding/json"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (u User) MarshalJSON() ([]byte, error) {
	r := struct {
		ID       int64  `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
	}

	return json.Marshal(&r)
}
