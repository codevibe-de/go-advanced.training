package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	setupRouter().Run("localhost:8080")
}

func getUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "123" {
		c.JSON(http.StatusOK, gin.H{"id": "123", "name": "John Doe"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users/:id", getUser)
	return r
}
