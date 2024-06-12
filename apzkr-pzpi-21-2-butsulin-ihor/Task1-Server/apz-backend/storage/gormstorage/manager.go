package gormstorage

import (
	"apz-backend/types/models"
	"gorm.io/gorm/clause"
)

func (g GormStorage) AddManager(manager models.Manager) error {
	db := g.db

	result := db.Create(&manager)

	return result.Error
}

func (g GormStorage) GetManagerByUserID(id uint) (*models.Manager, error) {
	db := g.db

	var manager models.Manager
	result := db.Model(&models.Manager{}).
		Preload(clause.Associations).
		Where(&models.Manager{UserID: id}).
		First(&manager)

	return &manager, result.Error
}

func (g GormStorage) GetManagers() ([]models.Manager, error) {
	db := g.db

	var managers []models.Manager
	result := db.Model(&models.Manager{}).
		Preload(clause.Associations).
		Find(&managers)

	return managers, result.Error
}
