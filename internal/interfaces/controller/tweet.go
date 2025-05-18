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

type ITweet interface {
	CreateTweet(ctx *gin.Context)
	GetTweetByID(ctx *gin.Context)
	UpdateTweet(ctx *gin.Context)
}

type Tweet struct {
	usecase usecase.ITweet
}

func NewTweet(usecase usecase.ITweet) Tweet {
	return Tweet{
		usecase: usecase,
	}
}

func (t Tweet) CreateTweet(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	var dtoTweet dto.TweetCreate
	err = json.Unmarshal(body, &dtoTweet)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	validationErr := pkg.ValidateStruct(dtoTweet)
	if validationErr != nil {
		apiErr := pkg.ToApiError(validationErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	if dtoTweet.UserID != ctx.GetInt64("userID") {
		apiErr := pkg.NewForbiddenApiError(pkg.NewForbiddenError("user not authorized", nil))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	tweet, usecaseErr := t.usecase.Create(dto.FromTweetCreateToTweet(dtoTweet))
	if usecaseErr != nil {
		apiErr := pkg.ToApiError(usecaseErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	ctx.JSON(http.StatusCreated, dto.FromTweetToTweetResponse(tweet))
}

func (t Tweet) GetTweetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidIDGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	tweet, usecaseErr := t.usecase.SearchByID(id)
	if usecaseErr != nil {
		apiErr := pkg.ToApiError(usecaseErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	ctx.JSON(http.StatusOK, dto.FromTweetToTweetResponse(tweet))
}

func (t Tweet) UpdateTweet(ctx *gin.Context) {
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

	var dtoTweet dto.TweetUpdate
	err = json.Unmarshal(body, &dtoTweet)
	if err != nil {
		apiErr := pkg.ToApiError(pkg.NewInvalidBodyGenericError(err))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	validationErr := pkg.ValidateStruct(dtoTweet)
	if validationErr != nil {
		apiErr := pkg.ToApiError(validationErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	if id != dtoTweet.ID {
		apiErr := pkg.NewForbiddenApiError(pkg.NewForbiddenError("tweet id mismatch", nil))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	if dtoTweet.UserID != ctx.GetInt64("userID") {
		apiErr := pkg.NewForbiddenApiError(pkg.NewForbiddenError("user not authorized", nil))
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	tweet, usecaseErr := t.usecase.Update(dto.FromTweetUpdateToTweet(dtoTweet))
	if usecaseErr != nil {
		apiErr := pkg.ToApiError(usecaseErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	ctx.JSON(http.StatusOK, dto.FromTweetToTweetResponse(tweet))
}
