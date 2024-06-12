package models

import (
	"gorm.io/gorm"
	"time"
)

type WeighingResult struct {
	gorm.Model
	SlotID uint      `json:"slotID"`
	Weight float64   `json:"weight"`
	Time   time.Time `json:"time"`
}
