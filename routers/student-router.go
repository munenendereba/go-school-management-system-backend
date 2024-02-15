package routers

import (
	"munenendereba/sms-backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeStudentRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/students", getStudents)

	return r
}

func StudentsRoutes(route *gin.Engine) {
	students := route.Group("/students")
	students.GET("", getStudents)
}

func getStudents(ctx *gin.Context) {
	res, err := controllers.GetStudents()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"students": res,
	})
}
