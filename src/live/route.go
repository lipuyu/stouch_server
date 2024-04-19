package live

import (
	"github.com/gin-gonic/gin"
	"stouch_server/src/live/controller"
)

func AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/focusUser", controller.PostFocusUser)
	rg.POST("/unfocusUser", controller.PostUnfocusUser)
}
