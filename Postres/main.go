package main

import (
	"Postres/controllers"
	"Postres/initializers"

	"Postres/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnecttoDB()
	initializers.DB.AutoMigrate(&models.Post{})
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
