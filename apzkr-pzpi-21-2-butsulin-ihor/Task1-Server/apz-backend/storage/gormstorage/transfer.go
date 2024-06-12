package gormstorage

import (
	"apz-backend/types/models"
	"gorm.io/gorm/clause"
)

func (g GormStorage) AddTransfer(transfer models.Transfer) error {
	db := g.db

	result := db.Create(&transfer)

	return result.Error
}

func (g GormStorage) GetTransfersByWarehouseID(warehouseID uint) ([]models.Transfer, error) {
	db := g.db

	var transfers []models.Transfer
	result := db.Model(&models.Transfer{}).
		Preload(clause.Associations).
		Preload("Car." + clause.Associations).
		Where(&models.Transfer{WarehouseID: warehouseID}).
		Find(&transfers)

	return transfers, result.Error
}

func (g GormStorage) GetTransfersByCarID(carID uint) ([]models.Transfer, error) {
	db := g.db

	var transfers []models.Transfer
	result := db.Model(&models.Transfer{}).
		Preload(clause.Associations).
		Where(&models.Transfer{CarID: carID}).
		Find(&transfers)

	return transfers, result.Error
}
