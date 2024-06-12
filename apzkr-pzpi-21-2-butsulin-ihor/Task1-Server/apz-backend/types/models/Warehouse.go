package models

import "gorm.io/gorm"

type Warehouse struct {
	gorm.Model
	StorageID uint     `json:"storageID"`
	Storage   Storage  `json:"storage"`
	Workers   []Worker `json:"workers"`
	Manager   Manager  `json:"manager"`
}
