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
	router.PUT("/recipes/:id", UpdateRecipeHandler)
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

func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var r Recipe
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index := -1
	for k, v := range recipes {
		if v.ID == id {
			index = k
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}
	recipes[index] = r
	c.JSON(http.StatusOK, r)
}
