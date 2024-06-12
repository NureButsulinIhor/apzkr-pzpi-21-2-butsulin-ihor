package task

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type Getter interface {
	GetTask(taskID uint) (*models.Task, error)
}

type Updater interface {
	UpdateTask(task *models.Task) error
}

type SlotUpdater interface {
	UpdateItemSlot(itemID uint, slotID uint) error
}

func SetDone(taskID uint, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.task.GetAll"),
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

	if user.Type != models.WorkerType {
		l.Debug("user is not manager")
		return errors.New("user is not manager")
	}

	l.Debug("getting worker from db")
	worker, err := cfg.Storage.GetWorkerByUserID(user.ID)
	if err != nil {
		l.Debug("err to get worker from db")
		return errors.New("no worker found")
	}

	l.Debug("getting task from db")
	task, err := cfg.Storage.GetTask(taskID)
	if err != nil {
		l.Debug("err to get task from db")
		return errors.New("no task found")
	}

	if task.WorkerID != worker.ID {
		l.Debug("worker is not owner of task")
		return errors.New("worker is not owner of task")
	}

	l.Debug("updating slots")
	fromSlot, err := cfg.Storage.GetSlot(task.FromSlotID)
	if err != nil {
		l.Error("err to get slot from db")
		return errors.New("internal error")
	}
	toSlot, err := cfg.Storage.GetSlot(task.ToSlotID)
	if err != nil {
		l.Error("err to get slot from db")
		return errors.New("internal error")
	}

	if fromSlot.Item == nil {
		l.Debug("item is not in slot")
		return errors.New("item is not in slot")
	}
	err = cfg.Storage.UpdateItemSlot(fromSlot.Item.ID, toSlot.ID)
	if err != nil {
		l.Error("err to update slot", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	l.Debug("updating task")
	task.Status = true
	err = cfg.Storage.UpdateTask(task)
	if err != nil {
		l.Error("err to update task", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	return nil
}
