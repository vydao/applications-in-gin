package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthcheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"env": app.config.env, "version": version})
}
