package registration

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"log/slog"
)

type DeviceAdder interface {
	AddDevice(device *models.Device) error
}

func RegisterDevice(slotID uint, cfg Configuration) (string, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.registration.RegisterDevice"),
	)

	l.Debug("processing auth data")
	_, payload, err := jwtauth.FromContext(cfg.Context)
	if err != nil {
		l.Debug("err to parse jwt")
		return "", errors.New("invalid JWT token")
	}

	l.Debug("processing claims")
	user, err := models.NewUserFromClaims(payload)
	if err != nil {
		l.Debug("err to parse claims")
		return "", errors.New("invalid claims")
	}

	if user.Type != models.AdminType {
		l.Debug("user is not admin")
		return "", errors.New("user is not admin")
	}

	l.Debug("creating new device")
	device := &models.Device{
		ID:     uuid.New(),
		SlotID: slotID,
	}
	err = cfg.Storage.AddDevice(device)
	if err != nil {
		l.Error("err to add device", slog.String("error", err.Error()))
		return "", errors.New("internal error")
	}

	l.Debug("generating token")
	_, tokenString, err := cfg.JWTAuth.Encode(device.GetClaims())
	if err != nil {
		l.Error("err to generate token", slog.String("error", err.Error()))
		return "", errors.New("internal error")
	}

	return tokenString, nil
}
