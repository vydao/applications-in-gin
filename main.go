// Recipes API
//
// This is a sample recipes API.
//
//	Schemes: http
//	Host: localhost:8000
//	BasePath: /
//	Version: 1.0.0
//
// Consumes:
//   - application/json
//
// Produces:
//   - application/json
//
// swagger:meta
package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
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
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	router.GET("/recipes/search", SearchRecipeHandler)
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

// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
// - application/json
// responses:
// '200':
// description: Successful operation
func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

// swagger:operation PUT /recipes/{id} recipes updateRecipe
// Update an existing recipe
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the recipe
//     required: true
//     type: string
//
// produces:
// - application/json
// responses:
// '200':
// description: Successful operation
// '400':
// description: Invalid input
// '404':
// description: Invalid recipe ID
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

// swagger:operation DELETE /recipes/{id} recipes deleteRecipe
// Delete an existing recipe
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the recipe
//     required: true
//     type: string
//
// produces:
// - application/json
// responses:
// '200':
// description: Successful operation
// '400':
// description: Invalid input
// '404':
// description: Invalid recipe ID
func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}
	recipes = append(recipes[:index], recipes[index+1:]...)
	c.JSON(http.StatusOK, gin.H{"msg": "Recipe has been deleted"})
}

// swagger:operation GET /recipes/search recipes searchRecipe
// Search recipes by tag or keyword
// ---
// parameters:
//   - name: tag
//     in: query
//     description: keyword to search
//     required: true
//     type: string
//
// produces:
// - application/json
// responses:
// '200':
// description: Successful operation
func SearchRecipeHandler(c *gin.Context) {
	tag := c.Query("tag")
	var result []Recipe

	for i := 0; i < len(recipes); i++ {
		found := false
		for _, t := range recipes[i].Tags {
			if strings.EqualFold(t, tag) {
				found = true
			}
		}
		if found {
			result = append(result, recipes[i])
		}
	}
	c.JSON(http.StatusOK, result)
}
