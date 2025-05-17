package service

import (
	"twitter-uala/internal/domain"
	"twitter-uala/internal/helpers"
	"twitter-uala/internal/repository"
	"twitter-uala/pkg"
)

type Auth struct {
	userRepository repository.IUser
}

func NewAuth(userRepository repository.IUser) Auth {
	return Auth{
		userRepository: userRepository,
	}
}

type AuthToken struct {
	Token string `json:"token"`
}

func (a Auth) Create(user domain.User) (AuthToken, pkg.Error) {
	var result AuthToken

	dbUser, err := a.userRepository.SelectByEmail(user.Email)
	if err != nil {
		return result, err
	}

	if !pkg.CheckHashedString(dbUser.Password, user.Password) {
		return result, pkg.NewGenericError("invalid password", nil)
	}

	token, generationErr := helpers.GenerateJwtToken(dbUser)
	if generationErr != nil {
		return result, pkg.NewFatalError("could not generate JWY token", generationErr)
	}

	result.Token = token

	return result, err
}
