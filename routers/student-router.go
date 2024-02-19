package routers

import (
	"munenendereba/sms-backend/controllers"
	"munenendereba/sms-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StudentsRoutes(route *gin.Engine) {
	students := route.Group("/students")
	students.GET("/", getStudents)
	students.GET("/:admissionNumber", getStudent)
	students.POST("/", createStudent)
	students.PUT("/", updateStudent)
	students.DELETE("/:admissionNumber", deleteStudent)
}

func getStudents(ctx *gin.Context) {
	students, err := controllers.GetStudents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, students)
}

func getStudent(ctx *gin.Context) {
	admissionNumber := ctx.Param("admissionNumber")
	student, err := controllers.GetStudent(admissionNumber)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if student == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "student not found"})

		return
	}

	ctx.JSON(http.StatusOK, student)
}

func createStudent(ctx *gin.Context) {
	var student models.Student

	err := ctx.Bind(&student)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := controllers.CreateStudent(&student)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func updateStudent(ctx *gin.Context) {
	var student models.Student

	err := ctx.Bind(&student)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := controllers.UpdateStudent(&student)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if res != nil {
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "student not found"})
	}
}

func deleteStudent(ctx *gin.Context) {
	admissionNumber := ctx.Param("admissionNumber")

	found, err := controllers.DeleteStudent(admissionNumber)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	} else {
		if found < 1 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "student not found",
			})
		} else {
			ctx.JSON(http.StatusNoContent, gin.H{
				"message": "student deleted",
			})
		}
	}
}
