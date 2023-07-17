package controllers

import (
	"gin_jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type genreInput struct {
	Genre string `json:"genre" form:"genre"`
}

func GetAllGenre(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var genre []models.Genre
	db.Find(&genre)

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

func GetGenreById(c *gin.Context) { // Get model if exist
    var genre models.Genre

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": genre})
}

func GetMoviesByGenreId(c *gin.Context) { // Get model if exist
    var movies []models.Movie

    db := c.MustGet("db").(*gorm.DB)

    if err := db.Where("genre_id = ?", c.Param("id")).Find(&movies).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": movies})
}

func CreateGenre(c *gin.Context) {
    // Validate input
    var input genreInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create Rating
    genre := models.Genre{Genre: input.Genre}
    db := c.MustGet("db").(*gorm.DB)
    db.Create(&genre)

    c.JSON(http.StatusOK, gin.H{"data": genre})
}

func UpdateGenre(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    // Get model if exist
    var genre models.Genre
    if err := db.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    // Validate input
    var input genreInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var updatedInput models.Genre
    updatedInput.Genre = input.Genre
    
    db.Model(&genre).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"data": genre})
}

func DeleteGenre(c *gin.Context) {
    // Get model if exist
    db := c.MustGet("db").(*gorm.DB)
    var genre models.Genre
    if err := db.Where("id = ?", c.Param("id")).First(&genre).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    db.Delete(&genre)

    c.JSON(http.StatusOK, gin.H{"data": "deleted succesfully"})
}