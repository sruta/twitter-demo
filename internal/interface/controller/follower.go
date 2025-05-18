package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"twitter-uala/internal/interface/dto"
	"twitter-uala/internal/usecase"
	"twitter-uala/pkg"
)

type IFollower interface {
	CreateFollower(ctx *gin.Context)
}

type Follower struct {
	usecase usecase.IFollower
}

func NewFollower(usecase usecase.IFollower) Follower {
	return Follower{
		usecase: usecase,
	}
}

func (f Follower) CreateFollower(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	var dtoFollower dto.FollowerCreate
	err = json.Unmarshal(body, &dtoFollower)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	validationErr := pkg.ValidateStruct(dtoFollower)
	if validationErr != nil {
		apiErr := pkg.ToApiError(validationErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	if dtoFollower.FollowerID != ctx.GetInt64("userID") {
		apiErr := pkg.NewForbiddenApiError(pkg.NewForbiddenError("user not authorized", nil))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	follower, usecaseErr := f.usecase.Create(dto.FromFollowerCreateToFollower(dtoFollower))
	if usecaseErr != nil {
		apiErr := pkg.ToApiError(usecaseErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	ctx.JSON(http.StatusCreated, dto.FromFollowerToFollowerResponse(follower))
}
