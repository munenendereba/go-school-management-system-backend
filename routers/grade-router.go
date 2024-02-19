package routers

import (
	"munenendereba/sms-backend/controllers"
	"munenendereba/sms-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GradeRoutes(route *gin.Engine) {
	gradeRoutes := route.Group("/grades")

	gradeRoutes.GET("/", getGrades)
	gradeRoutes.GET("/:id", getGrade)
	gradeRoutes.POST("/", createGrade)
	gradeRoutes.PUT("/", updateGrade)
	gradeRoutes.DELETE("/", deleteGrade)
}

func getGrades(ctx *gin.Context) {
	grades, err := controllers.GetGrades()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, grades)
}

func getGrade(ctx *gin.Context) {
	id := ctx.Param("id")

	grade, err := controllers.GetGrade(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if grade == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "grade not found"})

		return
	}

	ctx.JSON(http.StatusOK, grade)
}

func createGrade(ctx *gin.Context) {
	var grade models.Grade

	err := ctx.Bind(&grade)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	res, err := controllers.CreateGrade(&grade)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func updateGrade(ctx *gin.Context) {
	var grade models.Grade

	err := ctx.Bind(&grade)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)

		return
	}

	res, err := controllers.UpdateGrade(&grade)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)

		return
	}

	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "grade not found"})

		return
	}

	ctx.JSON(http.StatusOK, res)
}

func deleteGrade(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := controllers.DeleteGrade(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)

		return
	}

	if res < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "grade not found"})

		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"message": "grade deleted"})
}
