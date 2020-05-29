package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/joscelynjames/go-tutorial/models"
	"github.com/joscelynjames/go-tutorial/controllers"
)

func main() {
	// Define the router 
	// gin.Default gives you helpful middleware 
	r := gin.Default()

	db := models.SetupModels()

  r.Use(func(c *gin.Context) {
    c.Set("db", db)
    c.Next()
	})
	
	r.GET("/books", controllers.FindBooks)

	r.Run()
}