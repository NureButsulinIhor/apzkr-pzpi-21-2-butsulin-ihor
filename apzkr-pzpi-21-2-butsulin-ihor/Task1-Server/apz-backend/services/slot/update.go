package slot

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type Updater interface {
	UpdateSlot(slot *models.Slot) error
}

func Update(slotID uint, maxWeight float64, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.slot.Update"),
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

	l.Debug("getting slot from db")
	slot, err := cfg.Storage.GetSlot(slotID)
	if err != nil {
		l.Error("err to get slot from db")
		return errors.New("internal error")
	}

	if maxWeight <= 0 {
		l.Error("invalid maxWeight")
		return errors.New("invalid maxWeight")
	}

	l.Debug("updating slot from db")
	slot.MaxWeight = maxWeight
	err = cfg.Storage.UpdateSlot(slot)
	if err != nil {
		l.Error("err to delete slot from db")
		return errors.New("internal error")
	}

	return nil
}
