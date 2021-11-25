package appconf

import (
	"github.com/gin-gonic/gin"
	"stouch_server/src/appconf/controller"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.GET("/cdn", controller.GetCdn)
}
