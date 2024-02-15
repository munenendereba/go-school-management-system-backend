package routers

import "github.com/gin-gonic/gin"

func InitializeRouter() *gin.Engine {
	router := gin.Default()

	StudentsRoutes(router)

	return router
}
