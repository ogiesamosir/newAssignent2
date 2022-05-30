package models

import "time"

type CreateOrders struct {
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Item         []Items   `json:"items"`
}
