package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"twitter-demo/internal"
	"twitter-demo/internal/middleware"
)

func initRouter(c *internal.Container) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	api.POST("/user", c.UserController.CreateUser)
	api.POST("/login", c.AuthController.Login)

	secure := api.Group("/").Use(middleware.Auth())

	secure.GET("/user/:id", c.UserController.GetUserByID)
	secure.PUT("/user/:id", c.UserController.UpdateUser)

	secure.POST("/follower", c.FollowerController.CreateFollower)

	secure.POST("/tweet", c.TweetController.CreateTweet)
	secure.GET("/tweet/:id", c.TweetController.GetTweetByID)
	secure.PUT("/tweet/:id", c.TweetController.UpdateTweet)

	secure.GET("/timeline", c.TimelineController.GetTimeline)

	return r
}

func main() {
	c, err := internal.StartContainer()
	if err != nil {
		fmt.Println("error initializing container")
		return
	}

	r := initRouter(c)
	err = r.Run()
	if err != nil {
		fmt.Println("error starting server")
		return
	}
}
