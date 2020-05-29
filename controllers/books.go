package controllers

import (
	"github.com/gin-gonic/gin"
  "github.com/joscelynjames/go-tutorial/models"
)

func FindBooks(c *gin.Context) {
	db := c.MistGet("db").(*gorm.DB)

	var books []modles.Book
	db.find(&books)

	c.JSON(http.StatusOK, gin.H({"data": books }))
}