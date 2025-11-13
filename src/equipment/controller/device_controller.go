package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	authModel "stouch_server/src/auth/model"
	"stouch_server/src/common/er"
	"stouch_server/src/common/re"
	"stouch_server/src/core"
	"stouch_server/src/equipment/dto"
	"stouch_server/src/equipment/livepooleq"
	"stouch_server/src/equipment/model"
)

// PostByDeviceCode 绑定设备
func PostByDeviceCode(c *gin.Context) {
	deviceCode := c.Param("device_code")

	var device model.Device
	if has, err := core.Orm.Where("device_code = ?", deviceCode).Get(&device); err != nil {
		core.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, re.ErrorStr(err))
		return
	} else if !has {
		c.JSON(http.StatusOK, re.Error(er.EquipmentNotExistError))
		return
	}
	if device.UserId != 0 {
		c.JSON(http.StatusOK, re.Error(er.EquipmentAlreadyBindError))
		return
	}

	device.UserId = c.MustGet("user").(authModel.User).Id
	if _, err := core.Orm.Update(&device); err != nil {
		core.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, re.ErrorStr(err))
		return
	}
	c.JSON(http.StatusOK, re.Data(device))
}

// GetDeviceByUser 获取用户绑定的设备
func GetDeviceByUser(c *gin.Context) {
	var devices []model.Device
	userId := c.MustGet("user").(authModel.User).Id

	if err := core.Orm.Where("user_id = ?", userId).Find(&devices); err != nil {
		c.JSON(http.StatusInternalServerError, re.ErrorStr(err))
		return
	}

	c.JSON(http.StatusOK, re.Data(gin.H{"devices": devices}))
}

// SendCommand 发送命令到设备
func SendCommand(c *gin.Context) {
	// 绑定json
	var deviceCommandDTO dto.DeviceCommandDTO
	if err := c.BindJSON(&deviceCommandDTO); err != nil {
		c.JSON(http.StatusOK, re.Error(er.ParamsError))
		return
	}
	connMap := livepooleq.GetEquipmentConnMap()
	if conn, ok := connMap.Load(deviceCommandDTO.DeviceCode); ok {
		if err := conn.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(deviceCommandDTO.Command)); err != nil {
			c.JSON(http.StatusInternalServerError, re.ErrorStr(err))
			return
		}
		c.JSON(http.StatusOK, re.Data(""))
	} else {
		c.JSON(http.StatusNotFound, re.Error(er.EquipmentOfflineError))
	}
}

/**
上传土壤湿度
*/
func Post() {

}

/**
返回设备功能列表
*/
func GetDeviceFunctions() {

}
