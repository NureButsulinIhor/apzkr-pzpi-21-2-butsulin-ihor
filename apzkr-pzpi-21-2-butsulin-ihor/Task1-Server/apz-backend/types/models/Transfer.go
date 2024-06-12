package models

import (
	"gorm.io/gorm"
	"time"
)

type Transfer struct {
	gorm.Model
	CarID       uint      `json:"carID"`
	Car         Car       `json:"car"`
	WarehouseID uint      `json:"warehouseID"`
	Warehouse   Warehouse `json:"warehouse"`
	InDate      time.Time `json:"inDate"`
	OutDate     time.Time `json:"outDate"`
}
