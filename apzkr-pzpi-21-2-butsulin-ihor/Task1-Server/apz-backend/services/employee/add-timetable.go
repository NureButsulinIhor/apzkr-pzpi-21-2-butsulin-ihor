package employee

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
	"time"
)

type ManagerGetter interface {
	GetManagerByUserID(id uint) (*models.Manager, error)
}

type WorkerGetter interface {
	GetWorker(id uint) (*models.Worker, error)
}

type TimetableAdder interface {
	AddTimetable(timetable models.Timetable) error
}

func AddToTimetable(workerID uint, startWorkShift time.Duration, endWorkShift time.Duration, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.employee.AddToTimetable"),
	)

	l.Debug("validate start and end work shift")
	if startWorkShift >= endWorkShift {
		l.Debug("start work shift is greater or equal than end work shift")
		return errors.New("start work shift is greater or equal than end work shift")
	}

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

	l.Debug("getting worker from db")
	worker, err := cfg.Storage.GetWorker(workerID)
	if err != nil {
		l.Debug("err to get warehouse from db")
		return errors.New("wrong slot id")
	}

	if worker.WarehouseID != manager.WarehouseID {
		l.Debug("manager is not owner of worker's warehouse")
		return errors.New("manager is not manager of worker's warehouse")
	}

	l.Debug("adding worker to timetable")
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	timetable := models.Timetable{
		WorkerID:  workerID,
		StartTime: today.Add(startWorkShift),
		EndTime:   today.Add(endWorkShift),
	}
	err = cfg.Storage.AddTimetable(timetable)
	if err != nil {
		l.Error("err to add timetable", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	return nil
}
