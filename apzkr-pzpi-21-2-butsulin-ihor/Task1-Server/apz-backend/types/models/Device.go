package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Device struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	SlotID    uint           `json:"slotID"`
}

func (d *Device) GetClaims() map[string]interface{} {
	return map[string]interface{}{"id": d.ID.String()}
}

func NewDeviceFromClaims(claims map[string]interface{}) (*Device, error) {
	idString, ok := claims["id"].(string)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, errors.New("invalid uuid")
	}

	return &Device{
		ID: id,
	}, nil
}
