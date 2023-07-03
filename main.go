package main

import (
	"github.com/Gagas17ez/GoTest/controllers/productcontroller"
	"github.com/Gagas17ez/GoTest/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDB()

	r.GET("api/product", productcontroller.Index)
	r.GET("api/product/:id", productcontroller.Show)
	r.POST("api/product", productcontroller.Create)
	r.PUT("api/product/:id", productcontroller.Update)
	r.DELETE("api/product/", productcontroller.Delete)

	r.Run()
}
