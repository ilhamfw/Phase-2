package utils

import "github.com/gin-gonic/gin"

type ApiError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func ErrorMessage(c *gin.Context, apiError ApiError) {
    c.Abort()
    c.JSON(apiError.Code, gin.H{"error": apiError.Message})
}
