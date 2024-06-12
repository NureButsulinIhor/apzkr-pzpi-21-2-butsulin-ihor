package registration

import (
	"context"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type Storage interface {
	DeviceAdder
	UserAdder
	WorkerAdder
	ManagersGetter
	UsersByTypeGetter
	DeviceDeleter
	WorkerDeleter
}

type Configuration struct {
	Storage Storage
	Logger  *slog.Logger
	Context context.Context
	JWTAuth *jwtauth.JWTAuth
}
