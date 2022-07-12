package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaimeiherrera/go-gorm-example/db"
	"github.com/jaimeiherrera/go-gorm-example/models"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.User{}, models.Task{})

	r := gin.Default()
	r.GET("/")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
