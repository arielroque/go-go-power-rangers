package main

import (
	"example.com/go-go-power-rangers/controllers"
	"example.com/go-go-power-rangers/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/seasons", controllers.CreateSeason)
	router.GET("/seasons", controllers.FindSeasons)

	router.Run()

}
