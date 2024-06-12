package models

import "gorm.io/gorm"

type StorageType string

const (
	WarehouseType StorageType = "warehouse"
	CarType       StorageType = "car"
)

type Storage struct {
	gorm.Model
	Slots []Slot      `json:"slots"`
	Type  StorageType `json:"type"`
}
