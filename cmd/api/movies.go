package main

import (
	"net/http"
)

// Create new movie via "POST /v1/movies" endpoint
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// c.JSON(http.StatusCreated, nil)
}

// Get movie details by id "GET /v1/movies/:id"
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// id := c.Param("id")
	// c.JSON(http.StatusOK, gin.H{"id": id})
}
