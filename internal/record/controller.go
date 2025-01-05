package record

import "github.com/gin-gonic/gin"

func getHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "server is running!",
	})
}
