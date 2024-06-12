package gormstorage

import (
	"apz-backend/types/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (g GormStorage) AddWorker(worker models.Worker) error {
	db := g.db

	result := db.Create(&worker)

	return result.Error
}

func (g GormStorage) GetWorker(id uint) (*models.Worker, error) {
	db := g.db

	var worker models.Worker
	result := db.Model(&models.Worker{}).
		Preload(clause.Associations).First(&worker, id)

	return &worker, result.Error
}

func (g GormStorage) GetWorkerByUserID(id uint) (*models.Worker, error) {
	db := g.db

	var worker models.Worker
	result := db.Model(&models.Worker{}).
		Preload(clause.Associations).
		Where(&models.Worker{UserID: id}).
		First(&worker)

	return &worker, result.Error
}

func (g GormStorage) DeleteWorker(workerID uint) error {
	db := g.db

	result := db.Delete(&models.Worker{Model: gorm.Model{ID: workerID}})

	return result.Error
}

func (g GormStorage) AddTimetable(timetable models.Timetable) error {
	db := g.db

	result := db.Create(&timetable)

	return result.Error
}
