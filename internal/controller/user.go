package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/mail"
	"strconv"
	"twitter-uala/internal/domain"
	"twitter-uala/internal/service"
	"twitter-uala/pkg"
)

type User struct {
	service service.IUser
}

func NewUser(s service.IUser) User {
	return User{
		service: s,
	}
}

func (u User) GetUsers(ctx *gin.Context) {
	users, serviceErr := u.service.Search()
	if serviceErr != nil {
		apiErr := pkg.ToApiError(serviceErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u User) GetUserByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidIDGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	user, serviceErr := u.service.SearchByID(id)
	if serviceErr != nil {
		apiErr := pkg.ToApiError(serviceErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u User) PostUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	var user domain.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	validationErr := validateUser(user)
	if validationErr != nil {
		apiErr := pkg.ToApiError(validationErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	user, serviceErr := u.service.Create(user)
	if serviceErr != nil {
		apiErr := pkg.ToApiError(serviceErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func validateUser(user domain.User) pkg.Error {
	if user.Email == "" {
		return pkg.NewGenericError("user's email can't be empty", nil)
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return pkg.NewGenericError("user's email must be valid", err)
	}

	if user.Password == "" {
		return pkg.NewGenericError("user's password can't be empty", nil)
	}

	if len(user.Password) < 3 {
		return pkg.NewGenericError("user's password length can't be lower than three", nil)
	}

	if user.Username == "" {
		return pkg.NewGenericError("user's username can't be empty", nil)
	}

	return nil
}
