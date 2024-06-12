package models

import (
	"gorm.io/gorm"
	"time"
)

type Worker struct {
	gorm.Model
	UserID      uint        `json:"userID"`
	User        User        `json:"user"`
	WarehouseID uint        `json:"warehouseID"`
	Timetables  []Timetable `json:"timetables"`
}

func (w Worker) IsWorking() bool {
	for _, timetable := range w.Timetables {
		if timetable.StartTime.Before(time.Now()) && timetable.EndTime.After(time.Now()) {
			return true
		}
	}
	return false
}
