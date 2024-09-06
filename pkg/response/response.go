package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

func Error(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "error",
		"message": err,
	})
}
