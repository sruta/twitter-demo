package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"twitter-uala/internal"
	"twitter-uala/internal/middleware"
)

func initRouter(c *internal.Container) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	api.POST("/user", c.UserController.PostUser)
	api.POST("/login", c.AuthController.PostLogin)

	secure := api.Group("/").Use(middleware.Auth())
	//secure.PUT("/user/:id", c.UserController.UpdateUser)
	secure.GET("/user/:id", c.UserController.GetUserByID)
	//secure.GET("/user/:id/followers", c.UserController.GetUserFollowers)
	//secure.GET("/user/:id/follows", c.UserController.GetUserFollows)

	//secure.POST("/tweet", c.TweetController.PostTweet)
	//secure.PUT("/tweet/:id", c.TweetController.UpdateTweet)
	//secure.GET("/tweet/:id", c.TweetController.GetTweetByID)
	//secure.GET("/tweet/user/:userID", c.TweetController.GetTweetsByUserID)

	//secure.GET("/timeline", c.TimelineController.GetTimeline)

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
