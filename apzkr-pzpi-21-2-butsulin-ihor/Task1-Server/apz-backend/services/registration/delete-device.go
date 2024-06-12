package registration

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"log/slog"
)

type DeviceDeleter interface {
	DeleteDevice(deviceID uuid.UUID) error
}

func DeleteDevice(deviceID uuid.UUID, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.registration.DeleteDevice"),
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

	l.Debug("deleting device")
	err = cfg.Storage.DeleteDevice(deviceID)
	if err != nil {
		l.Debug("err to delete device", slog.String("error", err.Error()))
		return errors.New("no device found")
	}

	return nil
}
