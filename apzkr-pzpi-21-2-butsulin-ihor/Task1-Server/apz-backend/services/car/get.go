package car

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type Getter interface {
	GetCar(carID uint) (*models.Car, error)
}

func Get(carID uint, cfg Configuration) (*models.Car, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.car.Get"),
	)

	l.Debug("processing auth data")
	_, payload, err := jwtauth.FromContext(cfg.Context)
	if err != nil {
		l.Debug("err to parse jwt")
		return nil, errors.New("invalid JWT token")
	}

	l.Debug("processing claims")
	user, err := models.NewUserFromClaims(payload)
	if err != nil {
		l.Debug("err to parse claims")
		return nil, errors.New("invalid claims")
	}

	if user.Type != models.AdminType && user.Type != models.ManagerType {
		l.Debug("user is not admin")
		return nil, errors.New("user is not admin")
	}

	l.Debug("getting car")
	car, err := cfg.Storage.GetCar(carID)
	if err != nil {
		l.Debug("err to get car", slog.String("error", err.Error()))
		return nil, errors.New("no car found")
	}

	return car, nil
}
