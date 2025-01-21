package main

import (
	"github.com/alvaro-veiga/categories-microservice/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(route *gin.Engine) {
	categoryRoutes := route.Group("/categories")

	categoryRoutes.POST("/", controllers.CreateCategory)
}