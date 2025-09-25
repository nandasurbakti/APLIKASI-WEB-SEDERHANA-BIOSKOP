package routers

import (
	"bioskop_app/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", controllers.CreateBioskop)
	router.GET("/bioskop/AllBioskop", controllers.GetBioskop)
	router.PUT("/bioskop", controllers.UpdateBioskop)
	router.DELETE("/bioskop", controllers.DeleteBioskop)

	return router
}