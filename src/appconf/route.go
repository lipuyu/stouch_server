package appconf

import (
	"stouch_server/src/appconf/controller"

	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.GET("", controller.Get)
}
