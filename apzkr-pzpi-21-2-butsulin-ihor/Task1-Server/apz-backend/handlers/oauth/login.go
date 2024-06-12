package oauth

import (
	"apz-backend/services/login"
	"apz-backend/types"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type LoginData struct {
	GoogleJWT string `json:"googleJWT"`
}

type Response struct {
	JWT string `json:"jwt"`
}

func Login(logger *slog.Logger, jwtAuth *jwtauth.JWTAuth, googleClientID string, userGetter login.UserGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.oauth.Login"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody LoginData
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

		outputJWT, err := login.Google(requestBody.GoogleJWT, googleClientID, login.Configuration{
			Logger:     l,
			UserGetter: userGetter,
			Context:    r.Context(),
			JWTAuth:    jwtAuth,
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
