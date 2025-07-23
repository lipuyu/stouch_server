package test

import (
	"github.com/gin-gonic/gin"
	"stouch_server/src/websock/livepool"
)

func Hign2(c *gin.Context) {
	livepool.SendStringToAll("high2")
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func Low2(c *gin.Context) {
	livepool.SendStringToAll("low2")
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
