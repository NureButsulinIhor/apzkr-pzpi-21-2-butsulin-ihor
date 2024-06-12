package transfer

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

type CarGetter interface {
	GetCar(id uint) (*models.Car, error)
}

type Adder interface {
	AddTransfer(transfer models.Transfer) error
}

func Add(outDate time.Time, carID uint, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.transfer.Add"),
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

	l.Debug("getting car from db")
	_, err = cfg.Storage.GetCar(carID)
	if err != nil {
		l.Error("err to get warehouse from db")
		return errors.New("internal error")
	}

	inDate := time.Now()

	if inDate.After(outDate) || inDate.Equal(outDate) {
		l.Debug("inDate must be before outDate")
		return errors.New("inDate must be before outDate")
	}

	l.Debug("check if car is in warehouse")
	transfers, err := cfg.Storage.GetTransfersByCarID(carID)
	if err != nil {
		l.Error("err to get transfers from db", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	for _, transfer := range transfers {
		if transfer.InDate.Before(outDate) && transfer.OutDate.After(inDate) {
			l.Debug("car is in another warehouse")
			return errors.New("car is in another warehouse")
		}
	}

	transfer := models.Transfer{
		CarID:       carID,
		WarehouseID: manager.WarehouseID,
		InDate:      inDate,
		OutDate:     outDate,
	}
	err = cfg.Storage.AddTransfer(transfer)
	if err != nil {
		l.Error("err to add transfer", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	return nil
}
