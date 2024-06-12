package gormstorage

import "apz-backend/types/models"

func (g GormStorage) AddStorage(storage models.Storage) (uint, error) {
	db := g.db

	result := db.Create(&storage)

	return storage.ID, result.Error
}
