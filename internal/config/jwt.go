package config

import "time"

type JWT struct {
	Secret     string
	Expiration time.Duration
}

var JWTProd = JWT{
	Secret:     "A_SUPER_SECRET_STRING",
	Expiration: 24 * 60 * time.Minute,
}
