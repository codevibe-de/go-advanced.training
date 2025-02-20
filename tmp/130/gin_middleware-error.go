package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(errorHandler)
	r.GET("/error/advanced", advancedError)
	r.Run(":8080")
}

func advancedError(ctx *gin.Context) {
	_, err := businessLogic()
	if err != nil {
		ctx.Error(err)
	}
}

func errorHandler(c *gin.Context) {
	c.Next()
	c.JSON(
		http.StatusInternalServerError,
		gin.H{"message": fmt.Sprintf("Got %d error(s)", len(c.Errors))},
	)
}
