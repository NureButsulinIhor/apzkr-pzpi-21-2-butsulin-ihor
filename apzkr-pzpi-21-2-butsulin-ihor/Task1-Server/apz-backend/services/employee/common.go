package employee

import (
	"context"
	"log/slog"
)

type Storage interface {
	ManagerGetter
	WorkerGetter
	TimetableAdder
}

type Configuration struct {
	Storage Storage
	Logger  *slog.Logger
	Context context.Context
}
