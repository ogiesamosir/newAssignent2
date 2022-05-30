package models

import "time"

type Orders struct {
	ID           uint      `gorm:"primaryKey"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Item         []Items   `gorm:"foreignKey:OrderId"`
}
