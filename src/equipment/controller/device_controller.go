package controller

import (
	"errors"
	"net/http"
	authModel "stouch_server/src/auth/model"
	"stouch_server/src/common/er"
	"stouch_server/src/common/re"
	"stouch_server/src/core"
	"stouch_server/src/equipment/dto"
	"stouch_server/src/equipment/livepooleq"
	"stouch_server/src/equipment/model"
	"stouch_server/src/equipment/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

// 备土壤湿度、温度和开关信息
func GetDeviceInfo(c *gin.Context) {
	deviceCode := c.Param("deviceCode")

	var device model.Device
	if has, err := core.Orm.Where("device_code = ?", deviceCode).Get(&device); err != nil {
		core.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, re.ErrorStr(err))
		return
	} else if !has {
		c.JSON(http.StatusOK, re.Error(er.EquipmentNotExistError))
		return
	}

	flowerDevice, err := service.GetOrCreateFlowerDeviceInfo(deviceCode)
	if err != nil {
		core.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, re.ErrorStr(err))
		return
	}

	c.JSON(http.StatusOK, re.Data(gin.H{
		"humidity":    "30",
		"minHumidity": flowerDevice.MinHumidity,
		"maxHumidity": flowerDevice.MaxHumidity,
		"temperature": "30",
	}))
}

// SetDeviceMinHumidity 设置设备最小湿度
func SetDeviceMinHumidity(c *gin.Context) {
	deviceCode := c.Param("deviceCode")
	var body struct {
		MinHumidity *float64 `json:"minHumidity"`
	}
	if err := c.BindJSON(&body); err != nil || body.MinHumidity == nil {
		c.JSON(http.StatusOK, re.Error(er.ParamsError))
		return
	}
	if *body.MinHumidity < 0 || *body.MinHumidity > 100 {
		c.JSON(http.StatusOK, re.ErrorStr(errors.New("minHumidity invalid")))
		return
	}

	flowerDevice, err := service.GetOrCreateFlowerDeviceInfo(deviceCode)
	if err != nil {
		core.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, re.ErrorStr(err))
		return
	}

	if flowerDevice.MaxHumidity > 0 && *body.MinHumidity >= flowerDevice.MaxHumidity {
		c.JSON(http.StatusOK, re.ErrorStr(errors.New("minHumidity must be < maxHumidity")))
		return
	}

	flowerDevice.MinHumidity = *body.MinHumidity
	if _, err := core.Orm.Where("device_code = ?", deviceCode).Cols("min_humidity").Update(&flowerDevice); err != nil {
		core.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, re.ErrorStr(err))
		return
	}

	c.JSON(http.StatusOK, re.Data(gin.H{
		"deviceCode":  deviceCode,
		"minHumidity": flowerDevice.MinHumidity,
		"maxHumidity": flowerDevice.MaxHumidity,
	}))
}

// SetDeviceMaxHumidity 设置设备最大湿度
func SetDeviceMaxHumidity(c *gin.Context) {
	deviceCode := c.Param("deviceCode")
	var body struct {
		MaxHumidity *float64 `json:"maxHumidity"`
	}
	if err := c.BindJSON(&body); err != nil || body.MaxHumidity == nil {
		c.JSON(http.StatusOK, re.Error(er.ParamsError))
		return
	}
	if *body.MaxHumidity < 0 || *body.MaxHumidity > 100 {
		c.JSON(http.StatusOK, re.ErrorStr(errors.New("maxHumidity invalid")))
		return
	}

	flowerDevice, err := service.GetOrCreateFlowerDeviceInfo(deviceCode)
	if err != nil {
		core.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, re.ErrorStr(err))
		return
	}

	if flowerDevice.MinHumidity > 0 && *body.MaxHumidity <= flowerDevice.MinHumidity {
		c.JSON(http.StatusOK, re.ErrorStr(errors.New("maxHumidity must be > minHumidity")))
		return
	}

	flowerDevice.MaxHumidity = *body.MaxHumidity
	if _, err := core.Orm.Where("device_code = ?", deviceCode).Cols("max_humidity").Update(&flowerDevice); err != nil {
		core.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, re.ErrorStr(err))
		return
	}

	c.JSON(http.StatusOK, re.Data(gin.H{
		"deviceCode":  deviceCode,
		"minHumidity": flowerDevice.MinHumidity,
		"maxHumidity": flowerDevice.MaxHumidity,
	}))
}

// OpenDevice 打开设备
func OpenDevice(c *gin.Context) {
	deviceCode := c.Param("deviceCode")
	c.JSON(http.StatusOK, re.Data(gin.H{"deviceCode": deviceCode, "status": "ON"}))
}

// CloseDevice 关闭设备
func CloseDevice(c *gin.Context) {
	deviceCode := c.Param("deviceCode")
	c.JSON(http.StatusOK, re.Data(gin.H{
		"deviceCode": deviceCode,
		"status":     "OFF",
	}))
}
