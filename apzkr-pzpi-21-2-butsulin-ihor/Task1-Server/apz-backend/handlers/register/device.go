package register

import (
	"apz-backend/services/registration"
	"apz-backend/types"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type DeviceRequestBody struct {
	SlotID uint `json:"slotID"`
}

type Response struct {
	JWT string `json:"jwt"`
}

func Device(logger *slog.Logger, jwtAuth *jwtauth.JWTAuth, storage registration.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.register.Device"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody DeviceRequestBody
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

		outputJWT, err := registration.RegisterDevice(requestBody.SlotID, registration.Configuration{
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

		render.JSON(w, r, types.Response[Response]{
			Status: true,
			Body:   Response{JWT: outputJWT},
		})
	}
}
