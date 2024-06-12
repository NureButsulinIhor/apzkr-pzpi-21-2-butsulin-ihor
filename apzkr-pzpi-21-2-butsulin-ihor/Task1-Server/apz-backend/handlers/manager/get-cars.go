package manager

import (
	"apz-backend/services/car"
	"apz-backend/types"
	"apz-backend/types/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type GetCarsResponse struct {
	Cars []models.Car `json:"cars"`
}

func GetCars(logger *slog.Logger, storage car.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.manager.GetCars"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		cars, err := car.GetAll(car.Configuration{
			Logger:  l,
			Storage: storage,
			Context: r.Context(),
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

		render.JSON(w, r, types.Response[GetCarsResponse]{
			Status: true,
			Body:   GetCarsResponse{Cars: cars},
		})
	}
}
