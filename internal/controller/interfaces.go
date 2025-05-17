package controller

import "github.com/gin-gonic/gin"

type IUser interface {
	GetUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	PostUser(ctx *gin.Context)
}

type IAuth interface {
	PostLogin(ctx *gin.Context)
}
