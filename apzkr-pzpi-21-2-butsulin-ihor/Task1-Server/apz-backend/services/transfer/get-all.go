package transfer

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type ByCarGetter interface {
	GetTransfersByCarID(carID uint) ([]models.Transfer, error)
}

func GetAll(carID uint, cfg Configuration) ([]models.Transfer, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.transfer.GetAll"),
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

	if user.Type != models.ManagerType {
		l.Debug("user is not manager")
		return nil, errors.New("user is not manager")
	}

	l.Debug("getting transfers from db")
	transfers, err := cfg.Storage.GetTransfersByCarID(user.ID)
	if err != nil {
		l.Debug("err to get transfers from db")
		return nil, errors.New("no car found")
	}

	return transfers, nil
}
