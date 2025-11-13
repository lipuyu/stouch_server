package equipment

import (
	"github.com/gin-gonic/gin"
	"stouch_server/src/equipment/controller"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.GET("/:device_code/bind", controller.PostByDeviceCode)
	rg.GET("/device", controller.GetDeviceByUser)
	rg.POST("/command", controller.SendCommand)
}
