package slot

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

func Get(slotID uint, cfg Configuration) (*models.Slot, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.slot.Get"),
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

	if user.Type != models.AdminType && user.Type != models.ManagerType {
		l.Debug("user is not admin")
		return nil, errors.New("user is not admin")
	}

	l.Debug("getting slot from db")
	slot, err := cfg.Storage.GetSlot(slotID)
	if err != nil {
		l.Error("err to get slot from db")
		return nil, errors.New("internal error")
	}

	return slot, nil
}
