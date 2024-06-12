package device

import (
	"context"
	"log/slog"
)

type Storage interface {
	Getter
	WeighingResultSetter
	ManagerGetter
	WarehouseGetter
	SlotGetter
}

type Configuration struct {
	Storage Storage
	Logger  *slog.Logger
	Context context.Context
}
