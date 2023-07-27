package main

import (
	"github.com/bakhtiyorovdilshod/go-crud/initializers"
	"github.com/bakhtiyorovdilshod/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
