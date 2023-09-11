package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"published_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

var recipes []Recipe

func main() {
	router := gin.Default()
	router.POST("/recipes", CreateRecipeHandler)
	router.Run("localhost:8000")
}

func CreateRecipeHandler(c *gin.Context) {
	var r Recipe
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
		return
	}
	r.ID = xid.New().String()
}
