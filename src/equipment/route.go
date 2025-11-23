package equipment

import (
	"stouch_server/src/equipment/controller"

	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/:deviceCode/bind", controller.PostByDeviceCode)
	rg.GET("/:deviceCode/info", controller.GetDeviceInfo)
	rg.POST("/:deviceCode/maxHumidity", controller.SetDeviceMaxHumidity)
	rg.POST("/:deviceCode/minHumidity", controller.SetDeviceMinHumidity)
	rg.GET("/device", controller.GetDeviceByUser)
	rg.POST("/command", controller.SendCommand)
	rg.POST("/:deviceCode/on", controller.OpenDevice)
	rg.POST("/:deviceCode/off", controller.CloseDevice)
}
