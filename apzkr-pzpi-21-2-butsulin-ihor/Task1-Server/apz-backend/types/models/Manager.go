package models

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	WarehouseID uint `json:"warehouseID"`
	UserID      uint `json:"userID"`
	User        User `json:"user"`
}
