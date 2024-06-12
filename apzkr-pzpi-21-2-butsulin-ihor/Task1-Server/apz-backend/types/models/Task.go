package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	WorkerID   uint   `json:"workerID"`
	Worker     Worker `json:"worker"`
	FromSlotID uint   `json:"fromSlotID"`
	FromSlot   Slot   `json:"fromSlot"`
	ToSlotID   uint   `json:"toSlotID"`
	ToSlot     Slot   `json:"toSlot"`
	Status     bool   `json:"status"`
}
