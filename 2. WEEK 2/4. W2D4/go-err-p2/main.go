package main

import (
	"github.com/gin-gonic/gin"
	"go-err-p2/middleware"
	"go-err-p2/utils"
)

func main() {
	r := gin.Default()
	r.Use(middleware.HandleError())
	r.GET("/ping", func(c *gin.Context) {
		//panic("something went wrong")

		utils.SuccessMessage(c, &utils.SuccessAddData)
	})

	r.Run(":8080")

}
