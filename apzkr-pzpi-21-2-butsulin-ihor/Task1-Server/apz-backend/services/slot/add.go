package slot

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type Adder interface {
	AddSlot(slot models.Slot) error
}

func Add(storageID uint, maxWeight float64, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.slot.Add"),
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

	if user.Type != models.AdminType {
		l.Debug("user is not admin")
		return errors.New("user is not admin")
	}

	slot := models.Slot{
		StorageID: storageID,
		MaxWeight: maxWeight,
	}
	err = cfg.Storage.AddSlot(slot)
	if err != nil {
		l.Error("err to add slot", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	return nil
}
