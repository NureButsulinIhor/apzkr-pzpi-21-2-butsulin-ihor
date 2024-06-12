package item

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type Updater interface {
	UpdateItem(itemID uint, name, description string, weight float64) error
}

type Getter interface {
	GetItem(id uint) (*models.Item, error)
}

func Update(itemID uint, name, description string, weight float64, cfg Configuration) error {
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

	l.Debug("getting item from db")
	_, err = cfg.Storage.GetItem(itemID)
	if err != nil {
		l.Error("err to get item from db")
		return errors.New("internal error")
	}

	if name == "" || description == "" {
		l.Error("invalid name, description or weight")
		return errors.New("invalid name, description or weight")
	}

	l.Debug("updating item in db")
	err = cfg.Storage.UpdateItem(itemID, name, description, weight)
	if err != nil {
		l.Error("err to update item from db")
		return errors.New("internal error")
	}

	return nil
}
