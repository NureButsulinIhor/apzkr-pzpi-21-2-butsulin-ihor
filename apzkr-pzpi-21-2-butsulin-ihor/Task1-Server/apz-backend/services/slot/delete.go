package slot

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type Deleter interface {
	DeleteSlot(slotID uint) error
}

type Getter interface {
	GetSlot(slotID uint) (*models.Slot, error)
}

func Delete(slotID uint, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.slot.Delete"),
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

	if slot.Item != nil {
		l.Error("slot is not empty")
		return errors.New("slot is not empty")
	}

	l.Debug("deleting slot from db")
	err = cfg.Storage.DeleteSlot(slotID)
	if err != nil {
		l.Error("err to delete slot from db")
		return errors.New("internal error")
	}

	return nil
}
