package car

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
	"time"
)

type TransfersByCarGetter interface {
	GetTransfersByWarehouseID(warehouseID uint) ([]models.Transfer, error)
}

type ManagerGetter interface {
	GetManagerByUserID(id uint) (*models.Manager, error)
}

func GetAll(cfg Configuration) ([]models.Car, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.car.GetAll"),
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

	l.Debug("getting transfers from db")
	transfers, err := cfg.Storage.GetTransfersByWarehouseID(manager.WarehouseID)
	if err != nil {
		l.Error("err to get transfers from db", slog.String("error", err.Error()))
		return nil, errors.New("internal error")
	}

	var actualCars []models.Car
	for _, transfer := range transfers {
		if transfer.InDate.Before(time.Now()) && transfer.OutDate.After(time.Now()) {
			actualCars = append(actualCars, transfer.Car)
		}
	}

	return actualCars, nil
}
