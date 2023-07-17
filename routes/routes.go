package routes

import (
	controllers "gin_jwt/controllers"
	"gin_jwt/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    r := gin.Default()

    // set db to gin context
    r.Use(func(c *gin.Context) {
        c.Set("db", db)
    })

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    r.GET("/movies", controllers.GetAllMovie)
    r.GET("/movies/:id", controllers.GetMovieById)

    r.GET("/age-rating-categories", controllers.GetAllRating)
    r.GET("/age-rating-categories/:id", controllers.GetRatingById)
    r.GET("/age-rating-categories/:id/movies", controllers.GetMoviesByRatingId)

    r.GET("/genre", controllers.GetAllGenre)
    r.GET("/genre/:id", controllers.GetGenreById)
    r.GET("/genre/:id/movies", controllers.GetMoviesByGenreId)

    moviesMiddlewareRoute := r.Group("/movies")
    moviesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    moviesMiddlewareRoute.POST("/", controllers.CreateMovie)
    moviesMiddlewareRoute.PATCH("/:id", controllers.UpdateMovie)
    moviesMiddlewareRoute.DELETE("/:id", controllers.DeleteMovie)


    ratingMiddlewareRoute := r.Group("/age-rating-categories")
    ratingMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    ratingMiddlewareRoute.POST("/", controllers.CreateRating)
    ratingMiddlewareRoute.PATCH("/:id", controllers.UpdateRating)
    ratingMiddlewareRoute.DELETE("/:id", controllers.DeleteRating)

    genreMiddlewareRoute := r.Group("/genre")
    genreMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
    genreMiddlewareRoute.POST("/", controllers.CreateGenre)
    genreMiddlewareRoute.PATCH("/:id", controllers.UpdateGenre)
    genreMiddlewareRoute.DELETE("/:id", controllers.DeleteGenre)
    
    return r
}