package registration

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type WorkerAdder interface {
	AddWorker(worker models.Worker) error
}

type UserAdder interface {
	AddUser(user models.User) (uint, error)
}

func RegisterWorker(email string, warehouseID uint, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.registration.RegisterWorker"),
	)

	l.Debug("processing auth data")
	_, payload, err := jwtauth.FromContext(cfg.Context)
	if err != nil {
		l.Debug("err to parse jwt")
		return errors.New("invalid JWT token")
	}

	l.Debug("processing claims")
	user, err := models.NewUserFromClaims(payload)
	if err != nil {
		l.Debug("err to parse claims")
		return errors.New("invalid claims")
	}

	if user.Type != models.AdminType {
		l.Debug("user is not admin")
		return errors.New("user is not admin")
	}

	l.Debug("creating new user")
	newUser := models.User{
		Email: email,
		Type:  models.WorkerType,
	}
	userID, err := cfg.Storage.AddUser(newUser)
	if err != nil {
		l.Debug("err to add user", slog.String("error", err.Error()))
		return errors.New("user already exists")
	}

	l.Debug("creating new worker")
	worker := models.Worker{
		UserID:      userID,
		WarehouseID: warehouseID,
	}
	err = cfg.Storage.AddWorker(worker)
	if err != nil {
		l.Debug("err to add worker", slog.String("error", err.Error()))
		return errors.New("error to add worker, check warehouse id")
	}

	return nil
}
