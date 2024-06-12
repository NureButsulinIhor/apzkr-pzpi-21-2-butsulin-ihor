package device

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"log/slog"
)

type Connector interface {
	ConnectDevice(deviceID uuid.UUID, slotID uint) error
}

type ManagerGetter interface {
	GetManagerByUserID(id uint) (*models.Manager, error)
}

type WarehouseGetter interface {
	GetWarehouseBySlotID(id uint) (*models.Warehouse, error)
}

func Connect(slotID uint, deviceID uuid.UUID, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.device.Connect"),
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

	l.Debug("getting warehouse from db")
	warehouse, err := cfg.Storage.GetWarehouseBySlotID(slotID)
	if err != nil {
		l.Debug("err to get warehouse from db")
		return errors.New("wrong slot id")
	}

	if warehouse.ID != manager.WarehouseID {
		l.Debug("manager is not owner of warehouse")
		return errors.New("manager is not manager of warehouse")
	}

	l.Debug("connecting device")
	err = cfg.Storage.ConnectDevice(deviceID, slotID)
	if err != nil {
		l.Error("err to connect device", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	return nil
}
