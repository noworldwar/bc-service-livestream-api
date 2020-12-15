package utils

import (
	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, code int, errorMsg string) {
	c.Set("ErrorMsg", errorMsg)
	c.JSON(200, gin.H{"success": false, "error_code": code})
}
