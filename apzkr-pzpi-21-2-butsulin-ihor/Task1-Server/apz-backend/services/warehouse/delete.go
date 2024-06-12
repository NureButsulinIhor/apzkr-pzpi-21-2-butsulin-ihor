package warehouse

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type Deleter interface {
	DeleteWarehouse(warehouseID uint) error
}

func Delete(warehouseID uint, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.warehouse.Delete"),
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

	l.Debug("deleting warehouse")
	err = cfg.Storage.DeleteWarehouse(warehouseID)
	if err != nil {
		l.Debug("err to delete manager", slog.String("error", err.Error()))
		return errors.New("no warehouse found")
	}

	return nil
}
