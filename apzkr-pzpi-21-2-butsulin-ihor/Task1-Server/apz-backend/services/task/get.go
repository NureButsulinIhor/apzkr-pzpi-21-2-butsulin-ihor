package task

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

func Get(taskID uint, cfg Configuration) (*models.Task, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.task.Delete"),
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

	if user.Type != models.ManagerType && user.Type != models.WorkerType {
		l.Debug("user is not manager")
		return nil, errors.New("user is not manager")
	}

	l.Debug("getting manager from db")
	manager, err := cfg.Storage.GetManagerByUserID(user.ID)
	if err != nil {
		l.Debug("err to get manager from db")
		return nil, errors.New("no manager found")
	}

	l.Debug("getting task from db")
	task, err := cfg.Storage.GetTask(taskID)
	if err != nil {
		l.Error("err to get task from db")
		return nil, errors.New("internal error")
	}

	if task.Worker.WarehouseID != manager.WarehouseID {
		l.Debug("task is not in manager's warehouse")
		return nil, errors.New("task is not in manager's warehouse")
	}

	return task, nil
}
