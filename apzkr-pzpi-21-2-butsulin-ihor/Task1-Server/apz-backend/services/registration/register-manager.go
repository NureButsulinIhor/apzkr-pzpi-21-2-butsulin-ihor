package registration

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

func RegisterManager(email string, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.registration.RegisterManager"),
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

	l.Debug("creating new user")
	newUser := models.User{
		Email: email,
		Type:  models.ManagerType,
	}
	_, err = cfg.Storage.AddUser(newUser)
	if err != nil {
		l.Debug("err to add user", slog.String("error", err.Error()))
		return errors.New("user already exists")
	}

	return nil
}
