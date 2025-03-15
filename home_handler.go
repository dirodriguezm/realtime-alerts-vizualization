package main

import (
	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	component := home()
	c.Status(200)
	err := component.Render(c.Request.Context(), c.Writer)
	if err != nil {
		c.Error(err)
		c.JSON(500, gin.H{
			"message": "Template rendering failed",
		})
	}
}
