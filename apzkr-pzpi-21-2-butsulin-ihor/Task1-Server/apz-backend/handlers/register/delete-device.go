package register

import (
	"apz-backend/services/registration"
	"apz-backend/types"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type DeleteDeviceRequestBody struct {
	DeviceID uuid.UUID `json:"deviceID"`
}

func DeleteDevice(logger *slog.Logger, jwtAuth *jwtauth.JWTAuth, storage registration.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.register.DeleteDevice"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody DeleteDeviceRequestBody
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

		err = registration.DeleteDevice(requestBody.DeviceID,
			registration.Configuration{
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

		render.JSON(w, r, types.Response[any]{
			Status: true,
			Body:   nil,
		})
	}
}
