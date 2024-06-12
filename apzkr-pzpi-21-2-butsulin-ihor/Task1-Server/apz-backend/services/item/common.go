package item

import (
	"context"
	"log/slog"
)

type Storage interface {
	Adder
	Getter
	SlotGetter
	Updater
}

type Configuration struct {
	Storage Storage
	Logger  *slog.Logger
	Context context.Context
}
