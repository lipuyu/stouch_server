package service

import (
	"stouch_server/src/core"
	"stouch_server/src/equipment/model"
)

// getOrCreateFlowerDeviceInfo 获取或创建花卉设备扩展信息
func GetOrCreateFlowerDeviceInfo(deviceCode string) (model.FlowerDeviceInfo, error) {
	var flowerDevice model.FlowerDeviceInfo
	if has, err := core.Orm.Where("device_code = ?", deviceCode).Get(&flowerDevice); err != nil {
		return flowerDevice, err
	} else if !has {
		flowerDevice = model.FlowerDeviceInfo{
			DeviceCode:  deviceCode,
			MinHumidity: 20,
			MaxHumidity: 80,
		}
		if _, err := core.Orm.Insert(&flowerDevice); err != nil {
			return flowerDevice, err
		}
	}
	return flowerDevice, nil
}
