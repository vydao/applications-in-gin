package main

import (
	"encoding/json"
	"net/http"
	"os"
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
}

var recipes []Recipe

func init() {
	data, _ := os.ReadFile("recipes.json")
	_ = json.Unmarshal(data, &recipes)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", CreateRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.Run("localhost:8000")
}

func CreateRecipeHandler(c *gin.Context) {
	var r Recipe
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r.ID = xid.New().String()
	r.PublishedAt = time.Now()
	recipes = append(recipes, r)
	c.JSON(http.StatusCreated, r)
}

func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}
