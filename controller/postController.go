package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"learngo.com/initializers"
	"learngo.com/models"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		log.Fatal("some thing wrong with error ❌")
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostRead(c *gin.Context) {
	var post []models.Post
	initializers.DB.Find(&post)

	c.JSON(200, gin.H{
		"message": "Post fetch successfully",
		"post":    post,
	})
}

func ReadSinglePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"message": "Post fetch successfully",
		"post":    post,
	})
}

func UpdateSinglePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	// Bind the incoming JSON/body and handle bind errors
	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}

	var post models.Post

	// Fetch the existing post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "post not found",
		})
		return
	}

	// Update fields
	if err := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "failed to update post",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "post updated successfully",
		"post":    post,
	})
}
