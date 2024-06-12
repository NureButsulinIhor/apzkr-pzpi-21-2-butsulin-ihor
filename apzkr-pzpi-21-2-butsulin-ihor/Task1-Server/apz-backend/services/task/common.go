package task

import (
	"context"
	"log/slog"
)

type Storage interface {
	ManagerGetter
	WorkerByUserIDGetter
	ByWorkerGetter
	Getter
	Deleter
	BySlotGetter
	Updater
	WarehouseGetter
	SlotGetter
	Adder
	TransfersByWarehouseGetter
	SlotUpdater
}

type Configuration struct {
	Storage Storage
	Logger  *slog.Logger
	Context context.Context
}
