package login

import (
	"apz-backend/types/models"
	"cloud.google.com/go/auth/credentials/idtoken"
	"context"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type UserGetter interface {
	GetUserByEmail(email string) (*models.User, error)
	UpdateUserData(user *models.User) error
}

type Configuration struct {
	UserGetter UserGetter
	Logger     *slog.Logger
	Context    context.Context
	JWTAuth    *jwtauth.JWTAuth
}

func Google(googleJWT string, clientID string, cfg Configuration) (string, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.login.Google"),
	)

	l.Debug("processing jwt from Google")
	payload, err := idtoken.Validate(cfg.Context, googleJWT, clientID)
	if err != nil {
		l.Debug("err to parse jwt from Google")
		return "", errors.New("invalid JWT token")
	}

	l.Debug("processing claims")
	claimsUser, err := models.NewUserFromClaims(payload.Claims)
	if err != nil {
		l.Debug("err to parse claims from Google")
		return "", errors.New("invalid JWT token")
	}

	l.Debug("getting user from db")
	user, err := cfg.UserGetter.GetUserByEmail(claimsUser.Email)
	if err != nil {
		l.Debug("err to get user from db")
		return "", errors.New("no user found")
	}

	l.Debug("updating user data")
	user.Name = claimsUser.Name
	user.Picture = claimsUser.Picture
	err = cfg.UserGetter.UpdateUserData(user)
	if err != nil {
		l.Error("err to update user data", slog.String("error", err.Error()))
		return "", errors.New("internal error")
	}

	l.Debug("generating token")
	_, tokenString, err := cfg.JWTAuth.Encode(user.GetClaims())
	if err != nil {
		l.Error("err to generate token", slog.String("error", err.Error()))
		return "", errors.New("internal error")
	}

	return tokenString, nil
}
