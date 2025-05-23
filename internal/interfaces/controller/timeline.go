package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"twitter-demo/internal/interfaces/dto"
	"twitter-demo/internal/usecase"
	"twitter-demo/pkg"
)

type ITimeline interface {
	GetTimeline(ctx *gin.Context)
}

type Timeline struct {
	usecase usecase.ITimeline
}

func NewTimeline(usecase usecase.ITimeline) Timeline {
	return Timeline{
		usecase: usecase,
	}
}

func (t Timeline) GetTimeline(ctx *gin.Context) {
	tweets, usecaseErr := t.usecase.GetTimeline(ctx.GetInt64("userID"))
	if usecaseErr != nil {
		apiErr := pkg.ToApiError(usecaseErr)
		ctx.JSON(apiErr.GetStatus(), apiErr.GetResponse())
		return
	}

	ctx.JSON(http.StatusOK, dto.FromTweetsToTweetsResponse(tweets))
}
