package routes

import (
	"github.com/Djuanzz/gorest/controller"
	"github.com/Djuanzz/gorest/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// --- SET DB TO GIN CONTEXT
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// --- PUBLIC ROUTES
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)

	r.GET("/movies", controller.GetAllMovies)
	r.GET("/movies/:id", controller.GetMovieId)

	r.GET("/age-rating-categories", controller.GetAllRating)
	r.GET("/age-rating-categories/:id", controller.GetRatingId)

	// --- MOVIE ROUTES
	movieMiddleware := r.Group("/movies")
	movieMiddleware.Use(middlewares.JwtAuthMiddleware())
	movieMiddleware.POST("/", controller.CreateMovie)
	movieMiddleware.PATCH("/:id", controller.UpdateMovie)
	movieMiddleware.DELETE("/:id", controller.DeleteMovie)

	// --- AGE RATING CATEGORY ROUTES
	ratingMovieMiddleware := r.Group("/age-rating-categories")
	ratingMovieMiddleware.Use(middlewares.JwtAuthMiddleware())
	ratingMovieMiddleware.POST("/", controller.CreateRating)
	ratingMovieMiddleware.PATCH("/:id", controller.UpdateRating)
	ratingMovieMiddleware.DELETE("/:id", controller.DeleteRating)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
