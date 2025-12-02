package main

import (
	"learngo.com/initializers"
	"learngo.com/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.DbConnect()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
