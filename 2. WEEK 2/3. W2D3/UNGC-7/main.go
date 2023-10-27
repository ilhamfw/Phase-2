package main

import (
	"log"
	"morag_ecommerce/config"
	"morag_ecommerce/entity"
	"morag_ecommerce/handler"
	"morag_ecommerce/docs"
    

	"github.com/gin-gonic/gin"
    swaggerfiles "github.com/swaggo/files"
 ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	db, err := config.InitDB()
	// Auto Migrate
	db.AutoMigrate(&entity.Store{})
	db.AutoMigrate(&entity.Product{})
	db.AutoMigrate(&entity.LoginData{})
	if err != nil {
		log.Fatal("Error test")
	}
	defer db.Close()

	// Endpoint untuk Store
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/users/register", handler.RegisterStore)
	r.POST("/users/login", handler.LoginStore)

	// Endpoint untuk Product
	r.POST("/products", handler.CreateProduct)
	r.GET("/products", handler.GetProducts)
	r.GET("/products/:id", handler.GetProductByID)
	r.PUT("/products/:id", handler.UpdateProduct)
	r.DELETE("/products/:id", handler.DeleteProduct)

	//
	docs.SwaggerInfo.BasePath = "/v1/item"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()

    // go install github.com/swaggo/swag/cmd/swag@latest
}

