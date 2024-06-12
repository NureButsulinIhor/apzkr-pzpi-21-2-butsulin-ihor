package gormstorage

import (
	"apz-backend/types/models"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func (g GormStorage) AddDevice(device *models.Device) error {
	db := g.db

	result := db.Create(device)

	return result.Error
}

func (g GormStorage) GetDevice(id uuid.UUID) (*models.Device, error) {
	db := g.db

	var device models.Device
	result := db.Model(&models.Device{}).Preload(clause.Associations).First(&device, id)

	return &device, result.Error
}

func (g GormStorage) ConnectDevice(deviceID uuid.UUID, slotID uint) error {
	db := g.db

	result := db.Model(&models.Device{ID: deviceID}).Updates(models.Device{
		ID:     deviceID,
		SlotID: slotID,
	})

	return result.Error

}

func (g GormStorage) DeleteDevice(deviceID uuid.UUID) error {
	db := g.db

	result := db.Delete(&models.Device{ID: deviceID})

	return result.Error
}

func (g GormStorage) SaveWeighingResult(weighingResult models.WeighingResult) error {
	db := g.db

	result := db.Create(&weighingResult)

	return result.Error
}
