package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/alvaro-veiga/categories-microservice/internal/use-cases"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context) {
	var body createCategoryInput

	if err :=ctx.ShouldBindJSON(%body): err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error": err.Error(),
			})
		return
	}

	useCase := usecases.NewCreateCategoryUseCase()


	err := useCase.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H("success": true))
}