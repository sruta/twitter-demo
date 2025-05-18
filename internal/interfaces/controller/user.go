package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"twitter-uala/internal/interfaces/dto"
	"twitter-uala/internal/usecase"
	"twitter-uala/pkg"
)

type IUser interface {
	CreateUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

type User struct {
	usecase usecase.IUser
}

func NewUser(usecase usecase.IUser) User {
	return User{
		usecase: usecase,
	}
}

func (u User) CreateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	var dtoUser dto.UserCreate
	err = json.Unmarshal(body, &dtoUser)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	validationErr := pkg.ValidateStruct(dtoUser)
	if validationErr != nil {
		apiErr := pkg.ToApiError(validationErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	user, usecaseErr := u.usecase.Create(dto.FromUserCreateToUser(dtoUser))
	if usecaseErr != nil {
		apiErr := pkg.ToApiError(usecaseErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	ctx.JSON(http.StatusCreated, dto.FromUserToUserResponse(user))
}

func (u User) GetUsers(ctx *gin.Context) {
	users, usecaseErr := u.usecase.Search()
	if usecaseErr != nil {
		apiErr := pkg.ToApiError(usecaseErr)
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

	user, usecaseErr := u.usecase.SearchByID(id)
	if usecaseErr != nil {
		apiErr := pkg.ToApiError(usecaseErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	ctx.JSON(http.StatusOK, dto.FromUserToUserResponse(user))
}

func (u User) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidIDGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	var dtoUser dto.UserUpdate
	err = json.Unmarshal(body, &dtoUser)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	validationErr := pkg.ValidateStruct(dtoUser)
	if validationErr != nil {
		apiErr := pkg.ToApiError(validationErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	if id != dtoUser.ID {
		apiErr := pkg.NewForbiddenApiError(pkg.NewForbiddenError("user id mismatch", nil))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	if id != ctx.GetInt64("userID") {
		apiErr := pkg.NewForbiddenApiError(pkg.NewForbiddenError("user not authorized", nil))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	user, usecaseErr := u.usecase.Update(dto.FromUserUpdateToUser(dtoUser))
	if usecaseErr != nil {
		apiErr := pkg.ToApiError(usecaseErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	ctx.JSON(http.StatusOK, dto.FromUserToUserResponse(user))
}
