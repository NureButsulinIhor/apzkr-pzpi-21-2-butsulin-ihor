package register

import (
	"apz-backend/services/registration"
	"apz-backend/types"
	"apz-backend/types/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type GetManagerResponse struct {
	Managers      []models.Manager `json:"managers"`
	UnsetManagers []models.User    `json:"unsetManagers"`
}

func GetManagers(logger *slog.Logger, jwtAuth *jwtauth.JWTAuth, storage registration.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.register.GetManagers"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		managers, unsetManagers, err := registration.GetManagers(registration.Configuration{
			Logger:  l,
			Storage: storage,
			Context: r.Context(),
			JWTAuth: jwtAuth,
		})
		if err != nil {
			w.WriteHeader(400)
			render.JSON(w, r, types.Response[any]{
				Status: false,
				Error:  err.Error(),
				Body:   nil,
			})
			return
		}

		render.JSON(w, r, types.Response[GetManagerResponse]{
			Status: true,
			Body: GetManagerResponse{
				Managers:      managers,
				UnsetManagers: unsetManagers,
			},
		})
	}
}
