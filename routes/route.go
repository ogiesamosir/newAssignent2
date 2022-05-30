package routes

import (
	"newassignmen2/controllers"

	"github.com/gin-gonic/gin"
)

func StartRoute() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateItems)
	router.GET("/orders", controllers.GetItems)
	router.DELETE("/orders/:orderID", controllers.DeleteItems)
	router.PUT("/orders/:orderID", controllers.UpdateItems)

	return router

}
