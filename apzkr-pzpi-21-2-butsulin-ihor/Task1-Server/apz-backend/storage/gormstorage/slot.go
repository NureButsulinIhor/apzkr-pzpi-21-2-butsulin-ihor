package gormstorage

import (
	"apz-backend/types/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (g GormStorage) AddSlot(slot models.Slot) error {
	db := g.db

	result := db.Create(&slot)

	return result.Error
}

func (g GormStorage) GetSlot(id uint) (*models.Slot, error) {
	db := g.db

	var slot models.Slot
	result := db.Model(&models.Slot{}).Preload(clause.Associations).First(&slot, id)

	return &slot, result.Error
}

func (g GormStorage) DeleteSlot(slotID uint) error {
	db := g.db

	result := db.Delete(&models.Slot{Model: gorm.Model{ID: slotID}})

	return result.Error
}

func (g GormStorage) UpdateSlot(slot *models.Slot) error {
	db := g.db

	result := db.Model(&models.Slot{Model: gorm.Model{ID: slot.ID}}).Updates(models.Slot{
		Model:     gorm.Model{ID: slot.ID},
		MaxWeight: slot.MaxWeight,
	})

	return result.Error
}
