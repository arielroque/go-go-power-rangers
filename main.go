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
	router.GET("/seasons/:id", controllers.FindSeason)
	router.PATCH("/seasons/:id", controllers.UpdateSeason)
	router.DELETE("seasons/:id", controllers.DeleteSeason)

	error := router.Run()

	if error != nil {
		return
	}

}
