package warehouse

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type StorageAdder interface {
	AddStorage(storage models.Storage) (uint, error)
}

type Adder interface {
	AddWarehouse(warehouse models.Warehouse) (uint, error)
}

type UserGetter interface {
	GetUser(id uint) (*models.User, error)
}

type ManagerAdder interface {
	AddManager(manager models.Manager) error
}

func Add(managerUserID uint, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.warehouse.Add"),
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

	managerUser, err := cfg.Storage.GetUser(managerUserID)
	if err != nil {
		l.Debug("err to get user from db")
		return errors.New("no user found")
	}

	if managerUser.Type != models.ManagerType {
		l.Debug("user is not manager")
		return errors.New("user is not manager")
	}

	l.Debug("creating new storage")
	storage := models.Storage{
		Type: models.WarehouseType,
	}
	storageID, err := cfg.Storage.AddStorage(storage)
	if err != nil {
		l.Debug("err to add storage", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	l.Debug("creating new warehouse")
	warehouse := models.Warehouse{
		StorageID: storageID,
	}
	_, err = cfg.Storage.AddWarehouse(warehouse)
	if err != nil {
		l.Debug("err to add warehouse", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	l.Debug("creating new manager")
	manager := models.Manager{
		UserID:      managerUserID,
		WarehouseID: warehouse.ID,
	}
	err = cfg.Storage.AddManager(manager)
	if err != nil {
		l.Debug("err to add manager", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	return nil
}
