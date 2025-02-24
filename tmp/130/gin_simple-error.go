package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/error/simple", simpleError)
	r.Run(":8080")
}

func simpleError(ctx *gin.Context) {
	_, err := businessLogic()
	if err != nil {
		ctx.JSON(http.StatusTeapot, gin.H{"error": err.Error()})
	}
}
