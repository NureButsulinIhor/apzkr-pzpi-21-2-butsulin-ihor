package item

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type Adder interface {
	AddItem(Item models.Item) error
}

type SlotGetter interface {
	GetSlot(id uint) (*models.Slot, error)
}

func Add(name, description string, weight float64, slotID uint, cfg Configuration) error {
	l := cfg.Logger.With(
		slog.String("op", "services.item.Add"),
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
		l.Debug("err to get slot from db")
		return errors.New("no slot found")
	}

	if name == "" || description == "" || weight <= 0 {
		l.Error("invalid name, description or weight")
		return errors.New("invalid name, description or weight")
	}

	if slot.Item != nil || weight > slot.MaxWeight {
		l.Error("weight is more than maxWeight")
		return errors.New("weight is more than maxWeight")
	}

	item := models.Item{
		Name:        name,
		Description: description,
		Weight:      weight,
		SlotID:      slotID,
	}
	err = cfg.Storage.AddItem(item)
	if err != nil {
		l.Error("err to add item", slog.String("error", err.Error()))
		return errors.New("internal error")
	}

	return nil
}
