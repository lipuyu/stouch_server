package dto

type DeviceCommandDTO struct {
	DeviceCode string `json:"device_code" binding:"required"`
	Command    string `json:"command" binding:"required"`
}
