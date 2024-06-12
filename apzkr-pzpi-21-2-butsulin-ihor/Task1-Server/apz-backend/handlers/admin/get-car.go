package admin

import (
	"apz-backend/services/car"
	"apz-backend/types"
	"apz-backend/types/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

func GetCar(logger *slog.Logger, storage car.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.admin.GetCar"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		carIDStr := r.PathValue("carID")
		carID, err := strconv.ParseUint(carIDStr, 10, 64)
		if err != nil {
			w.WriteHeader(400)
			render.JSON(w, r, types.Response[any]{
				Status: false,
				Error:  "invalid request id",
				Body:   nil,
			})
			l.Debug("invalid request id")
			return
		}

		carModel, err := car.Get(uint(carID),
			car.Configuration{
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

		render.JSON(w, r, types.Response[models.Car]{
			Status: true,
			Body:   *carModel,
		})
	}
}
