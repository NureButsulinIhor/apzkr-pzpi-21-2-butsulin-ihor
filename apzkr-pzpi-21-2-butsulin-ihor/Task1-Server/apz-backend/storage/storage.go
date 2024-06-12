package storage

import (
	"apz-backend/services/car"
	"apz-backend/services/device"
	"apz-backend/services/employee"
	"apz-backend/services/item"
	"apz-backend/services/login"
	"apz-backend/services/registration"
	"apz-backend/services/slot"
	"apz-backend/services/task"
	"apz-backend/services/transfer"
	"apz-backend/services/warehouse"
)

type Storage interface {
	car.Storage
	device.Storage
	employee.Storage
	item.Storage
	login.UserGetter
	registration.Storage
	slot.Storage
	task.Storage
	transfer.Storage
	warehouse.Storage
}
