package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	StorageID uint    `json:"storageID"`
	Storage   Storage `json:"storage"`
	OwnerID   uint    `json:"ownerID"`
	Owner     User    `json:"owner"`
}
