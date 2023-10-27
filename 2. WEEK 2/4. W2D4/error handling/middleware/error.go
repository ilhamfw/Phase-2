package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func HandleError() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if r := recover(); r != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
                c.Abort()
            }
        }()
        c.Next()
    }
}
