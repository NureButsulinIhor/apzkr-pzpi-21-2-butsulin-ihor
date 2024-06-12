package slot

import (
	"context"
	"log/slog"
)

type Storage interface {
	Adder
	Getter
	Deleter
	Updater
}

type Configuration struct {
	Storage Storage
	Logger  *slog.Logger
	Context context.Context
}
