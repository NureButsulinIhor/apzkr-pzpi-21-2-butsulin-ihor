package models

import "gorm.io/gorm"

type Slot struct {
	gorm.Model
	MaxWeight       float64          `json:"maxWeight"`
	Item            *Item            `json:"item"`
	ItemID          *uint            `json:"itemID"`
	Device          *Device          `json:"device"`
	WeighingResults []WeighingResult `json:"weighingResults"`
	StorageID       uint             `json:"storageID"`
}
