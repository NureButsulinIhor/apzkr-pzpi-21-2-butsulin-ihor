package warehouse

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type WarehousesGetter interface {
	GetWarehouses() ([]models.Warehouse, error)
}

func GetAll(cfg Configuration) ([]models.Warehouse, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.warehouse.GetAll"),
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

	if user.Type != models.AdminType {
		l.Debug("user is not admin")
		return nil, errors.New("user is not admin")
	}

	warehouses, err := cfg.Storage.GetWarehouses()
	if err != nil {
		l.Debug("err to get warehouse from db")
		return nil, errors.New("no warehouse found")
	}

	return warehouses, nil
}
