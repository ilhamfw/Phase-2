package middleware

import (
	"github.com/gin-gonic/gin"
	"go-err-p2/utils"
)

func HandleError() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				utils.ErrorMessage(c, &utils.ErrInternalServer)
			}
		}()
		c.Next()

	}
}
