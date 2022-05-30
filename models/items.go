package models

type Items struct {
	ID          uint   `gorm:"primaryKey"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderId     uint   `json:"orderId"`
}
