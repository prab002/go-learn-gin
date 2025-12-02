package main

import (
	"github.com/gin-gonic/gin"
	"learngo.com/controller"
	"learngo.com/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.DbConnect()
}

func main() {
	r := gin.Default()

	// Group all post routes under /posts
	posts := r.Group("/posts")
	{
		posts.POST("/", controller.PostsCreate)        // POST /posts
		posts.GET("/", controller.PostRead)            // GET /posts
		posts.GET("/:id", controller.ReadSinglePost)   // GET /posts/:id
		posts.PUT("/:id", controller.UpdateSinglePost) // PUT /posts/:id
		// you can add:
		// posts.DELETE("/:id", controller.DeletePost)
	}

	r.Run()
}
