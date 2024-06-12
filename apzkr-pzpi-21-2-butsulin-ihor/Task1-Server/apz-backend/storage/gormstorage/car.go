package gormstorage

import (
	"apz-backend/types/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (g GormStorage) AddCar(car models.Car) (uint, error) {
	db := g.db

	result := db.Create(&car)

	return car.ID, result.Error
}

func (g GormStorage) DeleteCar(carID uint) error {
	db := g.db

	result := db.Delete(&models.Car{Model: gorm.Model{ID: carID}})

	return result.Error
}

func (g GormStorage) GetCar(carID uint) (*models.Car, error) {
	db := g.db

	var car models.Car
	result := db.Model(&models.Car{}).
		Preload(clause.Associations).
		Preload("Storage."+clause.Associations).
		Preload("Storage.Slots."+clause.Associations).
		Preload("Owner."+clause.Associations).
		First(&car, carID)

	return &car, result.Error
}

func (g GormStorage) GetCars() ([]models.Car, error) {
	db := g.db

	var cars []models.Car
	result := db.Model(&models.Car{}).
		Preload(clause.Associations).
		Preload("Owner." + clause.Associations).
		Find(&cars)

	return cars, result.Error
}
