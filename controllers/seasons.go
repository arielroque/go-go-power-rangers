package controllers

import (
	"net/http"

	"example.com/go-go-power-rangers/models"
	"github.com/gin-gonic/gin"
)

func DeleteSeason(c *gin.Context) {
	var season models.Season

	error := models.DB.Where("id = ?", c.Param("id")).First(&season).Error

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&season)
}

func UpdateSeason(c *gin.Context) {
	var season models.Season

	error := models.DB.Where("id = ?", c.Param("id")).First(&season).Error

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateSeasonInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&season).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": season})
}

func FindSeason(c *gin.Context) {
	var season models.Season

	err := models.DB.Where("id = ?", c.Param("id")).First(&season).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": season})
}

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
