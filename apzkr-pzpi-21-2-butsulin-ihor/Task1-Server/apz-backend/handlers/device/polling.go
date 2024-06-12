package device

import (
	"apz-backend/services/device"
	"apz-backend/types"
	"apz-backend/types/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"time"
)

type PollingRequestData struct {
	Weight float64   `json:"weight"`
	Time   time.Time `json:"time"`
}

type PollingResponseData struct {
	Weight float64 `json:"weight"`
}

func Polling(logger *slog.Logger, storage device.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.device.Polling"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody PollingRequestData
		err := render.DecodeJSON(r.Body, &requestBody)
		if err != nil {
			w.WriteHeader(400)
			render.JSON(w, r, types.Response[any]{
				Status: false,
				Error:  "invalid request body",
				Body:   nil,
			})
			l.Debug("err in decoding json")
			return
		}

		weight, err := device.Polling(models.WeighingResult{
			Weight: requestBody.Weight,
			Time:   requestBody.Time,
		}, device.Configuration{
			Storage: storage,
			Logger:  l,
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

		render.JSON(w, r, types.Response[PollingResponseData]{
			Status: true,
			Body:   PollingResponseData{Weight: weight},
		})
	}
}
