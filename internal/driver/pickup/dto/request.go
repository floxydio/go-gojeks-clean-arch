package dto

type OrderPickup struct {
	OrderId  string `json:"order_id" form:"order_id"`
	DriverId string `json:"driver_id" form:"driver_id"`
}
