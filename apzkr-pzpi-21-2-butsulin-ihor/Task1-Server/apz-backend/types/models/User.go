package models

import "gorm.io/gorm"

type UserType string

const (
	WorkerType  UserType = "worker"
	ManagerType UserType = "manager"
	AdminType   UserType = "admin"
)

type User struct {
	gorm.Model
	Email   string   `json:"email" gorm:"unique"`
	Name    string   `json:"name"`
	Picture string   `json:"picture"`
	Type    UserType `json:"type"`
}

func NewUserFromClaims(claims map[string]interface{}) (user User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	idFloat, _ := claims["id"].(float64)
	id := uint(idFloat)
	email := claims["email"].(string)
	name := claims["name"].(string)
	picture := claims["picture"].(string)
	userTypeString, _ := claims["type"].(string)
	userType := UserType(userTypeString)

	return User{
		Model:   gorm.Model{ID: id},
		Email:   email,
		Name:    name,
		Picture: picture,
		Type:    userType,
	}, err
}

func (u User) GetClaims() map[string]interface{} {
	return map[string]interface{}{
		"id":      float64(u.ID),
		"email":   u.Email,
		"name":    u.Name,
		"picture": u.Picture,
		"type":    string(u.Type),
	}
}
