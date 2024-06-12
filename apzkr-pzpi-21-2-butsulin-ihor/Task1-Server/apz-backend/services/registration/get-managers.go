package registration

import (
	"apz-backend/types/models"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"log/slog"
)

type ManagersGetter interface {
	GetManagers() ([]models.Manager, error)
}

type UsersByTypeGetter interface {
	GetUsersByType(usersType models.UserType) ([]models.User, error)
}

func GetManagers(cfg Configuration) ([]models.Manager, []models.User, error) {
	l := cfg.Logger.With(
		slog.String("op", "services.registration.GetManagers"),
	)

	l.Debug("processing auth data")
	_, payload, err := jwtauth.FromContext(cfg.Context)
	if err != nil {
		l.Debug("err to parse jwt")
		return nil, nil, errors.New("invalid JWT token")
	}

	l.Debug("processing claims")
	user, err := models.NewUserFromClaims(payload)
	if err != nil {
		l.Debug("err to parse claims")
		return nil, nil, errors.New("invalid claims")
	}

	if user.Type != models.AdminType {
		l.Debug("user is not admin")
		return nil, nil, errors.New("user is not admin")
	}

	l.Debug("getting managers from db")
	managers, err := cfg.Storage.GetManagers()
	if err != nil {
		l.Debug("err to get managers from db")
		return nil, nil, errors.New("internal error")
	}

	l.Debug("getting users from db")
	users, err := cfg.Storage.GetUsersByType(models.ManagerType)
	if err != nil {
		l.Debug("err to get users from db")
		return nil, nil, errors.New("internal error")
	}
	var unsettedManagers []models.User
firstLevelTag:
	for _, managerUser := range users {
		for _, manager := range managers {
			if managerUser.ID == manager.UserID {
				continue firstLevelTag
			}
		}

		unsettedManagers = append(unsettedManagers, managerUser)
	}

	return managers, unsettedManagers, nil
}
