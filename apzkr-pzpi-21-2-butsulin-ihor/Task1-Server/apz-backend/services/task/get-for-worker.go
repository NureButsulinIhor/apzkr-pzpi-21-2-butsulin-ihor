package task

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type ByWorkerGetter interface {
	GetTasksByWorkerID(workerID uint) ([]models.Task, error)
}

type WorkerByUserIDGetter interface {
	GetWorkerByUserID(id uint) (*models.Worker, error)
}

func GetForWorker(cfg Configuration) ([]models.Task, error) {
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

	if user.Type != models.WorkerType {
		l.Debug("user is not manager")
		return nil, errors.New("user is not manager")
	}

	l.Debug("getting worker from db")
	worker, err := cfg.Storage.GetWorkerByUserID(user.ID)
	if err != nil {
		l.Debug("err to get worker from db")
		return nil, errors.New("no worker found")
	}

	l.Debug("getting tasks from db")
	tasks, err := cfg.Storage.GetTasksByWorkerID(worker.ID)
	if err != nil {
		l.Error("err to get tasks from db")
		return nil, errors.New("internal error")
	}

	return tasks, nil
}
