package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "okay",
		})
	})

	r.Run("8080")
}
