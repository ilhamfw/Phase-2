package utils

import "github.com/gin-gonic/gin"

func SuccessMessage(c *gin.Context, apiError *APIError) *gin.Context {
	c.Abort()
	c.JSON(200, gin.H{"code": apiError.Code, "message": apiError.Message})
	return c
}
