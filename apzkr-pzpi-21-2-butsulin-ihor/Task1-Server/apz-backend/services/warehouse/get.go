package warehouse

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type Getter interface {
	GetWarehouse(warehouseID uint) (*models.Warehouse, error)
}

func Get(warehouseID uint, cfg Configuration) (*models.Warehouse, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.warehouse.Get"),
	)

	l.Debug("processing auth data")
	_, payload, err := jwtauth.FromContext(cfg.Context)
	if err != nil {
		l.Debug("err to parse jwt")
		return nil, errors.New("invalid JWT token")
	}

	l.Debug("processing claims")
	user, err := models.NewUserFromClaims(payload)
	if err != nil {
		l.Debug("err to parse claims")
		return nil, errors.New("invalid claims")
	}

	if user.Type != models.AdminType && user.Type != models.ManagerType {
		l.Debug("user is not admin")
		return nil, errors.New("user is not admin")
	}

	l.Debug("getting warehouse")
	warehouse, err := cfg.Storage.GetWarehouse(warehouseID)
	if err != nil {
		l.Debug("err to get car", slog.String("error", err.Error()))
		return nil, errors.New("no warehouse found")
	}

	return warehouse, nil
}
