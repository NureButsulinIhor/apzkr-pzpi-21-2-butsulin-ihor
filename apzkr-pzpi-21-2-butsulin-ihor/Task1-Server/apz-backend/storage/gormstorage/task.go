package gormstorage

import (
	"apz-backend/types/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (g GormStorage) AddTask(task models.Task) error {
	db := g.db

	result := db.Create(&task)

	return result.Error
}

func (g GormStorage) GetTasksByWorkerID(workerID uint) ([]models.Task, error) {
	db := g.db

	var tasks []models.Task
	result := db.Model(&models.Task{}).
		Preload(clause.Associations).
		Where(&models.Task{WorkerID: workerID}).
		Find(&tasks)

	return tasks, result.Error
}

func (g GormStorage) GetTask(taskID uint) (*models.Task, error) {
	db := g.db

	var task models.Task
	result := db.Model(&models.Task{}).
		Preload(clause.Associations).First(&task, taskID)

	return &task, result.Error
}

func (g GormStorage) GetTasksByFromSlotID(slotID uint) ([]models.Task, error) {
	db := g.db

	var tasks []models.Task
	result := db.Model(&models.Task{}).
		Preload(clause.Associations).
		Where(&models.Task{FromSlotID: slotID}).
		Find(&tasks)

	return tasks, result.Error
}

func (g GormStorage) GetTasksByToSlotID(slotID uint) ([]models.Task, error) {
	db := g.db

	var tasks []models.Task
	result := db.Model(&models.Task{}).
		Preload(clause.Associations).
		Where(&models.Task{ToSlotID: slotID}).
		Find(&tasks)

	return tasks, result.Error
}

func (g GormStorage) UpdateTask(task *models.Task) error {
	db := g.db

	result := db.Model(&models.Task{Model: gorm.Model{ID: task.ID}}).Updates(models.Task{
		Model:      gorm.Model{ID: task.ID},
		WorkerID:   task.WorkerID,
		FromSlotID: task.FromSlotID,
		ToSlotID:   task.ToSlotID,
		Status:     task.Status,
	})

	return result.Error
}

func (g GormStorage) DeleteTask(taskID uint) error {
	db := g.db

	result := db.Delete(&models.Task{Model: gorm.Model{ID: taskID}})

	return result.Error
}
