package task

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type Deleter interface {
	DeleteTask(taskID uint) error
}

func Delete(taskID uint, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.task.Delete"),
	)

	l.Debug("processing auth data")
	_, payload, err := jwtauth.FromContext(cfg.Context)
	if err != nil {
		l.Debug("err to parse jwt")
		return errors.New("invalid JWT token")
	}

	user, err := models.NewUserFromClaims(payload)
	if err != nil {
		l.Debug("err to parse claims")
		return errors.New("invalid claims")
	}

	if user.Type != models.ManagerType {
		l.Debug("user is not manager")
		return errors.New("user is not manager")
	}

	l.Debug("getting manager from db")
	manager, err := cfg.Storage.GetManagerByUserID(user.ID)
	if err != nil {
		l.Debug("err to get manager from db")
		return errors.New("no manager found")
	}

	l.Debug("getting task from db")
	task, err := cfg.Storage.GetTask(taskID)
	if err != nil {
		l.Error("err to get task from db")
		return errors.New("internal error")
	}

	if task.Worker.WarehouseID != manager.WarehouseID {
		l.Debug("task is not in manager's warehouse")
		return errors.New("task is not in manager's warehouse")
	}

	if task.Status {
		l.Debug("task is done")
		return errors.New("task is done")
	}

	l.Debug("deleting task from db")
	err = cfg.Storage.DeleteTask(taskID)
	if err != nil {
		l.Error("err to delete task from db")
		return errors.New("internal error")
	}

	return nil
}
