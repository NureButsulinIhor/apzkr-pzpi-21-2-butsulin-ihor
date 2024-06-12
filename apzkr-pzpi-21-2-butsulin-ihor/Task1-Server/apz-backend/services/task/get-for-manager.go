package task

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type ManagerGetter interface {
	GetManagerByUserID(id uint) (*models.Manager, error)
}

func GetForManager(cfg Configuration) ([]models.Task, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.task.GetAll"),
	)

	l.Debug("processing auth data")
	_, payload, err := jwtauth.FromContext(cfg.Context)
	if err != nil {
		l.Debug("err to parse jwt")
		return nil, errors.New("invalid JWT token")
	}

	user, err := models.NewUserFromClaims(payload)
	if err != nil {
		l.Debug("err to parse claims")
		return nil, errors.New("invalid claims")
	}

	if user.Type != models.ManagerType {
		l.Debug("user is not manager")
		return nil, errors.New("user is not manager")
	}

	l.Debug("getting manager from db")
	manager, err := cfg.Storage.GetManagerByUserID(user.ID)
	if err != nil {
		l.Debug("err to get manager from db")
		return nil, errors.New("no manager found")
	}

	warehouse, err := cfg.Storage.GetWarehouse(manager.WarehouseID)
	if err != nil {
		l.Debug("err to get warehouse from db")
		return nil, errors.New("no warehouse found")
	}

	var tasks []models.Task
	for _, worker := range warehouse.Workers {
		workerTasks, err := cfg.Storage.GetTasksByWorkerID(worker.ID)
		if err != nil {
			l.Debug("err to get tasks from db")
			continue
		}
		tasks = append(tasks, workerTasks...)
	}

	if len(tasks) == 0 {
		l.Debug("no tasks found")
		return nil, nil
	}

	return tasks, nil
}
