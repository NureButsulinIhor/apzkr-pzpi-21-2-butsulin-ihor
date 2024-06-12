package gormstorage

import (
	"apz-backend/types/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormStorage struct {
	db *gorm.DB
}

func NewStorage(connectionString string) (*GormStorage, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	err = initTable(db, &models.Item{}, err)
	err = initTable(db, &models.Device{}, err)
	err = initTable(db, &models.WeighingResult{}, err)
	err = initTable(db, &models.Slot{}, err)
	err = initTable(db, &models.Storage{}, err)
	err = initTable(db, &models.User{}, err)
	err = initTable(db, &models.Manager{}, err)
	err = initTable(db, &models.Timetable{}, err)
	err = initTable(db, &models.Car{}, err)
	err = initTable(db, &models.Task{}, err)
	err = initTable(db, &models.Worker{}, err)
	err = initTable(db, &models.Warehouse{}, err)
	err = initTable(db, &models.Transfer{}, err)

	return &GormStorage{db}, err
}

func initTable(db *gorm.DB, table any, prevError error) error {
	if prevError == nil {
		return db.AutoMigrate(table)
	}

	return prevError
}
