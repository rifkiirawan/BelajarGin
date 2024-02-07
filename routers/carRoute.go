package routers

import (
	"belajar-gin/controllers"
	// C:\Users\Rifki\go\pkg\mod\github.com\gin-gonic
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)

	router.PUT("/cars/:carID", controllers.UpdateCar)

	router.GET("/cars/:carID", controllers.GetCar)

	router.GET("/cars/allcars", controllers.GetAllCar)

	router.DELETE("/car/:carID", controllers.DeleteCar)

	return router
}
