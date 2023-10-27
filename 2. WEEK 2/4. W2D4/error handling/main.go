package main

import (
	"github.com/gin-gonic/gin"
	"handle_error/middleware"
	"handle_error/utils"
)

func main() {
	r := gin.Default()
	r.Use(middleware.HandleError())
	
	r.GET("/ping", func(c *gin.Context) {
		// Panic jika sesuatu berjalan salah
		// panic("something went wrong")
		
		// Menggunakan fungsi utils.SuccessMessage
		utils.SuccessMessage(c, utils.SuccessAddData)
	})

	r.Run(":8080")
}
