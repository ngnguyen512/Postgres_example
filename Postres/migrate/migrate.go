package main

import (
	"Postres/initializers"
	"Postres/models"
)

func init() {
	initializers.ConnecttoDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
