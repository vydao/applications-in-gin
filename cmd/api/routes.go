package main

import "github.com/gin-gonic/gin"

func (app *application) routes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/healthcheck", app.healthcheckHandler)
	v1.POST("/movies", app.createMovieHandler)
	v1.GET("/movies/:id", app.showMovieHandler)
	return r
}
