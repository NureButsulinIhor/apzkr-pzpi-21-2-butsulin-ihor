package warehouse

import (
	"context"
	"log/slog"
)

type Storage interface {
	ManagerGetter
	StorageAdder
	Adder
	UserGetter
	ManagerAdder
	Deleter
	Getter
	WarehousesGetter
}

type Configuration struct {
	Storage Storage
	Logger  *slog.Logger
	Context context.Context
}
