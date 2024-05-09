package controller

import (
	"net/http"
	"time"

	"github.com/Djuanzz/gorest/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type movieInput struct {
	Title               string `json:"title"`
	Year                int    `json:"year"`
	AgeRatingCategoryId int    `json:"age_rating_category_id"`
}

// CreateMovie godoc
// @Summary Create New Movie.
// @Description Creating a new Movie.
// @Tags Movie
// @Param Body body movieInput true "the body to create a new movie"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Movie
// @Router /movies [post]
func CreateMovie(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// --- VALIDATE
	var input movieInput
	var rating models.AgeRatingCategory

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.AgeRatingCategoryId).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Age rating category not found!"})
		return
	}

	// --- CREATE
	movie := models.Movie{Title: input.Title, Year: input.Year, AgeRatingCategoryId: input.AgeRatingCategoryId}
	db.Create(&movie)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// GetAllMovies godoc
// @Summary Get all movies.
// @Description Get a list of Movies.
// @Tags Movie
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.Movie
// @Router /movies [get]
func GetAllMovies(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var movies []models.Movie
	db.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// GetMovieById godoc
// @Summary Get Movie.
// @Description Get a Movie by id.
// @Tags Movie
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "movie id"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [get]
func GetMovieId(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var movie models.Movie
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// UpdateMovie godoc
// @Summary Update Movie.
// @Description Update movie by id.
// @Tags Movie
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "movie id"
// @Param Body body movieInput true "the body to update an movie"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [patch]
func UpdateMovie(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var movie models.Movie
	var rating models.AgeRatingCategory
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// --- VALIDATE
	var input movieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.AgeRatingCategoryId).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Age rating category not found!"})
		return
	}

	// --- UPDATE
	db.Model(&movie).Updates(models.Movie{Title: input.Title,
		Year:                input.Year,
		AgeRatingCategoryId: input.AgeRatingCategoryId,
		UpdatedAt:           time.Now()})

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// DeleteMovie godoc
// @Summary Delete one movie.
// @Description Delete a movie by id.
// @Tags Movie
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "movie id"
// @Success 200 {object} map[string]boolean
// @Router /movie/{id} [delete]
func DeleteMovie(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var movie models.Movie
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&movie)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
