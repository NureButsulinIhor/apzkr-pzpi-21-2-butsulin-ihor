package transfer

import (
	"context"
	"log/slog"
)

type Storage interface {
	ManagerGetter
	ByCarGetter
	CarGetter
	Adder
}

type Configuration struct {
	Storage Storage
	Logger  *slog.Logger
	Context context.Context
}
