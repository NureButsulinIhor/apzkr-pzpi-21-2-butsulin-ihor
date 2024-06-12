package manager

import (
	"apz-backend/services/device"
	"apz-backend/types"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type ConnectDeviceRequestData struct {
	DeviceID uuid.UUID `json:"deviceID"`
	SlotID   uint      `json:"slotID"`
}

func ConnectDevice(logger *slog.Logger, storage device.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.manager.ConnectDevice"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody ConnectDeviceRequestData
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

		err = device.Connect(requestBody.SlotID, requestBody.DeviceID, device.Configuration{
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

		render.JSON(w, r, types.Response[any]{
			Status: true,
			Body:   nil,
		})
	}
}
