package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"twitter-demo/internal/interfaces/dto"
	"twitter-demo/internal/usecase"
	"twitter-demo/pkg"
)

type IAuth interface {
	Login(ctx *gin.Context)
}

type Auth struct {
	service usecase.IAuth
}

func NewAuth(s usecase.IAuth) Auth {
	return Auth{
		service: s,
	}
}

func (a Auth) Login(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	var dtoCredentials dto.AuthLogin
	err = json.Unmarshal(body, &dtoCredentials)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	validationErr := pkg.ValidateStruct(dtoCredentials)
	if validationErr != nil {
		apiErr := pkg.ToApiError(validationErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	token, serviceErr := a.service.CreateToken(dtoCredentials.Email, dtoCredentials.Password)
	if serviceErr != nil {
		apiErr := pkg.ToApiError(serviceErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	result := dto.AuthLoginResponse{
		Token: token,
	}

	ctx.JSON(http.StatusCreated, result)
}
