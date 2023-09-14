package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create new movie via "POST /v1/movies" endpoint
func (app *application) createMovieHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

// Get movie details by id "GET /v1/movies/:id"
func (app *application) showMovieHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}
