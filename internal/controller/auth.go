package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"twitter-uala/internal/domain"
	"twitter-uala/internal/service"
	"twitter-uala/pkg"
)

type Auth struct {
	service service.IAuth
}

func NewAuth(s service.IAuth) Auth {
	return Auth{
		service: s,
	}
}

func (a Auth) PostLogin(ctx *gin.Context) {
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

	token, serviceErr := a.service.Create(user)
	if serviceErr != nil {
		apiErr := pkg.ToApiError(serviceErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	ctx.JSON(http.StatusCreated, token)
}
