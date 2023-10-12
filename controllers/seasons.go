package controllers

import (
	"net/http"

	"example.com/go-go-power-rangers/models"
	"github.com/gin-gonic/gin"
)

func FindSeasons(c *gin.Context) {
	var seasons []models.Season

	models.DB.Find(&seasons)

	c.JSON(http.StatusOK, gin.H{"data": seasons})
}

func CreateSeason(c *gin.Context) {
	var input models.CreateSeasonInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	season := models.Season{Title: input.Title, Description: input.Description}
	models.DB.Create(&season)

	c.JSON(http.StatusOK, gin.H{"data": season})

}
