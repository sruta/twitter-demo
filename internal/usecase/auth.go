package usecase

import (
	"twitter-uala/internal/infraestructure/repository"
	"twitter-uala/pkg"
)

type IAuth interface {
	CreateToken(email string, password string) (string, pkg.Error)
}

type Auth struct {
	userRepository repository.IUser
}

func NewAuth(userRepository repository.IUser) Auth {
	return Auth{
		userRepository: userRepository,
	}
}

func (a Auth) CreateToken(email string, password string) (string, pkg.Error) {
	var token string
	dbUser, err := a.userRepository.SelectByEmail(email)
	if err != nil {
		return token, err
	}

	if !pkg.CheckHashedString(dbUser.Password, password) {
		return token, pkg.NewGenericError("invalid password", nil)
	}

	token, generationErr := pkg.GenerateJwtToken(dbUser.ID)
	if generationErr != nil {
		return token, pkg.NewFatalError("could not generate JWT token", generationErr)
	}

	return token, nil
}
