package gormstorage

import (
	"apz-backend/types/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (g GormStorage) AddWarehouse(warehouse models.Warehouse) (uint, error) {
	db := g.db

	result := db.Create(&warehouse)

	return warehouse.ID, result.Error
}

func (g GormStorage) GetWarehouseBySlotID(id uint) (*models.Warehouse, error) {
	slot, err := g.GetSlot(id)
	if err != nil {
		return nil, err
	}

	db := g.db

	var warehouse models.Warehouse
	result := db.Model(&models.Warehouse{}).
		Preload(clause.Associations).
		Preload("Storage." + clause.Associations).
		Preload("Storage.Slots." + clause.Associations).
		Preload("Manager." + clause.Associations).
		Preload("Workers." + clause.Associations).
		Where(&models.Warehouse{StorageID: slot.StorageID}).
		First(&warehouse)

	return &warehouse, result.Error
}

func (g GormStorage) GetWarehouses() ([]models.Warehouse, error) {
	db := g.db

	var warehouses []models.Warehouse
	result := db.Model(&models.Warehouse{}).
		Preload(clause.Associations).
		Find(&warehouses)

	return warehouses, result.Error
}

func (g GormStorage) GetWarehouse(id uint) (*models.Warehouse, error) {
	db := g.db

	var warehouse models.Warehouse
	result := db.Model(&models.Warehouse{}).
		Preload(clause.Associations).
		Preload("Storage."+clause.Associations).
		Preload("Storage.Slots."+clause.Associations).
		Preload("Manager."+clause.Associations).
		Preload("Workers."+clause.Associations).
		First(&warehouse, id)

	return &warehouse, result.Error
}

func (g GormStorage) DeleteWarehouse(warehouseID uint) error {
	db := g.db

	result := db.Delete(&models.Warehouse{Model: gorm.Model{ID: warehouseID}})

	return result.Error
}
