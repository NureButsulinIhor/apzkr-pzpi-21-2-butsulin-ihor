package device

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"log/slog"
)

type ManagerGetter interface {
	GetManagerByUserID(id uint) (*models.Manager, error)
}

type WarehouseGetter interface {
	GetWarehouseBySlotID(id uint) (*models.Warehouse, error)
}

type Getter interface {
	GetDevice(id uuid.UUID) (*models.Device, error)
}

type WeighingResultSetter interface {
	SaveWeighingResult(weighingResult models.WeighingResult) error
}

type SlotGetter interface {
	GetSlot(id uint) (*models.Slot, error)
}

func Polling(weighingResult models.WeighingResult, cfg Configuration) (float64, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.device.Polling"),
	)

	l.Debug("processing jwt from device")
	_, payload, err := jwtauth.FromContext(cfg.Context)
	if err != nil {
		l.Debug("err to parse jwt from device")
		return 0, errors.New("invalid JWT token")
	}

	l.Debug("processing claims")
	claimsDevice, err := models.NewDeviceFromClaims(payload)
	if err != nil || claimsDevice.ID.String() == "" {
		l.Debug("err to parse claims from device")
		return 0, errors.New("invalid claims")
	}

	l.Debug("getting device from db")
	device, err := cfg.Storage.GetDevice(claimsDevice.ID)
	if err != nil {
		l.Debug("err to get device from db")
		return 0, errors.New("no device found")
	}

	weighingResult.SlotID = device.SlotID

	l.Debug("adding weighing result to db")
	err = cfg.Storage.SaveWeighingResult(weighingResult)
	if err != nil {
		l.Error("err to add weighing result")
		return 0, errors.New("internal error")
	}

	l.Debug("getting slot from db")
	slot, err := cfg.Storage.GetSlot(device.SlotID)
	if err != nil {
		l.Debug("err to get slot from db")
		return 0, errors.New("internal error")
	}

	var weight float64
	if slot.Item != nil {
		weight = slot.Item.Weight
	}

	return weight, nil
}
