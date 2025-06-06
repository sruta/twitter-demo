package usecase

import (
	"twitter-demo/internal/domain"
	"twitter-demo/internal/infraestructure/repository"
	"twitter-demo/pkg"
)

type IUser interface {
	Search() ([]domain.User, pkg.Error)
	SearchByID(id int64) (domain.User, pkg.Error)
	Create(user domain.User) (domain.User, pkg.Error)
	Update(user domain.User) (domain.User, pkg.Error)
}

type User struct {
	userRepository repository.IUser
}

func NewUser(userRepository repository.IUser) User {
	return User{
		userRepository: userRepository,
	}
}

func (u User) Search() ([]domain.User, pkg.Error) {
	return u.userRepository.Select()
}

func (u User) SearchByID(id int64) (domain.User, pkg.Error) {
	return u.userRepository.SelectByID(id)
}

func (u User) Create(user domain.User) (domain.User, pkg.Error) {
	repeatedUser, err := u.userRepository.SelectByEmail(user.Email)
	if err != nil && !pkg.IsNotFound(err) {
		return user, err
	}
	if repeatedUser.ID != 0 {
		return user, pkg.NewGenericError("user already exist with that email", nil)
	}

	repeatedUser, err = u.userRepository.SelectByUsername(user.Username)
	if err != nil && !pkg.IsNotFound(err) {
		return user, err
	}
	if repeatedUser.ID != 0 {
		return user, pkg.NewGenericError("user already exist with that username", nil)
	}

	hashedPassword, hashingErr := pkg.HashString(user.Password)
	if hashingErr != nil {
		return user, pkg.NewFatalError("password hashing failed", hashingErr)
	}
	user.Password = hashedPassword

	user, err = u.userRepository.Insert(user)
	if err != nil {
		return user, err
	}

	return u.userRepository.SelectByID(user.ID)
}

func (u User) Update(user domain.User) (domain.User, pkg.Error) {
	repeatedUser, err := u.userRepository.SelectByUsername(user.Username)
	if err != nil && !pkg.IsNotFound(err) {
		return user, err
	}
	if repeatedUser.ID != 0 && repeatedUser.ID != user.ID {
		return user, pkg.NewGenericError("user already exist with that username", nil)
	}

	dbUser, err := u.userRepository.SelectByID(user.ID)
	if err != nil {
		return user, err
	}

	dbUser.Username = user.Username

	dbUser, err = u.userRepository.Update(dbUser)
	if err != nil {
		return user, err
	}

	return u.userRepository.SelectByID(dbUser.ID)
}
