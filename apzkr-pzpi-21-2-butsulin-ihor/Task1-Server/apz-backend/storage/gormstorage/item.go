package gormstorage

import (
	"apz-backend/types/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (g GormStorage) AddItem(item models.Item) error {
	db := g.db

	result := db.Create(&item)

	return result.Error
}

func (g GormStorage) GetItem(id uint) (*models.Item, error) {
	db := g.db

	var item models.Item
	result := db.Model(&models.Item{}).Preload(clause.Associations).First(&item, id)

	return &item, result.Error
}

func (g GormStorage) UpdateItem(itemID uint, name, description string, weight float64) error {
	db := g.db

	result := db.Model(&models.Item{Model: gorm.Model{ID: itemID}}).Updates(models.Item{
		Model:       gorm.Model{ID: itemID},
		Name:        name,
		Description: description,
		Weight:      weight,
	})

	return result.Error
}

func (g GormStorage) UpdateItemSlot(itemID uint, slotID uint) error {
	db := g.db

	result := db.Model(&models.Item{Model: gorm.Model{ID: itemID}}).Updates(models.Item{
		SlotID: slotID,
	})

	return result.Error
}
