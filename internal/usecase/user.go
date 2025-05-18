package usecase

import (
	"twitter-uala/internal/domain"
	"twitter-uala/internal/infraestructure/repository"
	"twitter-uala/pkg"
)

type IUser interface {
	Search() ([]domain.User, pkg.Error)
	SearchByID(id int64) (domain.User, pkg.Error)
	Create(user domain.User) (domain.User, pkg.Error)
	Update(user domain.User) (domain.User, pkg.Error)
}

type User struct {
	rdb            pkg.DB
	userRepository repository.IUser
}

func NewUser(rdb pkg.DB, userRepository repository.IUser) User {
	return User{
		rdb:            rdb,
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

	return u.userRepository.Insert(user)
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

	return u.userRepository.Update(dbUser)
}
