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
	api.POST("/user", c.UserController.CreateUser) //OK
	api.POST("/login", c.AuthController.Login)     //OK

	secure := api.Group("/").Use(middleware.Auth())

	secure.GET("/user/:id", c.UserController.GetUserByID) //OK
	secure.PUT("/user/:id", c.UserController.UpdateUser)  //OK

	secure.POST("/follower", c.FollowerController.CreateFollower) //OK

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
