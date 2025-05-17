package service

import (
	"twitter-uala/internal/domain"
	"twitter-uala/pkg"
)

type IUser interface {
	Search() ([]domain.User, pkg.Error)
	SearchByID(id int64) (domain.User, pkg.Error)
	Create(user domain.User) (domain.User, pkg.Error)
}

type IAuth interface {
	Create(user domain.User) (AuthToken, pkg.Error)
}
