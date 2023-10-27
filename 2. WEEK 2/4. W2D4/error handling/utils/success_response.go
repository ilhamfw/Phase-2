package utils

import "github.com/gin-gonic/gin"

func SuccessMessage(c *gin.Context, apiError *APIError) *gin.Context {
    c.JSON(208, gin.H{
        "code":    apiError.Code,
        "message": apiError.Message,
    })

    return c
}
