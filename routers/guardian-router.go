package routers

import (
	"munenendereba/sms-backend/controllers"
	"munenendereba/sms-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GuardianRoutes(route *gin.Engine) {
	guardians := route.Group("/guardians")

	guardians.GET("/", getGuardians)
	guardians.GET("/:id", getGuardian)
	guardians.POST("/", createGuardian)
	guardians.PUT("/", updateGuardian)
	guardians.DELETE("/:id", deleteGuardian)
}

func getGuardians(ctx *gin.Context) {
	guardians, err := controllers.GetGuardians()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, guardians)
}

func getGuardian(ctx *gin.Context) {
	id := ctx.Param("id")

	guardian, err := controllers.GetGuardian(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
	}

	if guardian == nil {
		ctx.JSON(http.StatusNotFound,
			gin.H{"message": "guadian not found"})
	}

	ctx.JSON(http.StatusOK, guardian)
}

func createGuardian(ctx *gin.Context) {
	var guardian models.Guardian

	err := ctx.BindJSON(&guardian)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	res, err := controllers.CreateGuardian(&guardian)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func updateGuardian(ctx *gin.Context) {
	var guardian models.Guardian

	err := ctx.BindJSON(&guardian)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})

		return
	}

	res, err := controllers.UpdateGuardian(&guardian)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})

		return
	}

	if res != nil {
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "guardian not found"})
	}
}

func deleteGuardian(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := controllers.DeleteGuardian(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	if res < 1 {
		ctx.JSON(http.StatusNotFound,
			gin.H{"message": "guardian not found"})
	} else {
		ctx.JSON(http.StatusNoContent,
			gin.H{"message": "guardian deleted"})
	}
}
