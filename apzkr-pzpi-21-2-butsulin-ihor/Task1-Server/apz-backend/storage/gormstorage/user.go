package gormstorage

import (
	"apz-backend/types/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (g GormStorage) AddUser(user models.User) (uint, error) {
	db := g.db

	result := db.Create(&user)

	return user.ID, result.Error
}

func (g GormStorage) GetUser(id uint) (*models.User, error) {
	db := g.db

	var user models.User
	result := db.Model(&models.User{}).
		Preload(clause.Associations).First(&user, id)

	return &user, result.Error
}

func (g GormStorage) GetUserByEmail(email string) (*models.User, error) {
	db := g.db

	var user models.User
	result := db.Model(&models.User{}).
		Preload(clause.Associations).
		Where(&models.User{Email: email}).
		First(&user)

	return &user, result.Error
}

func (g GormStorage) GetUsersByType(usersType models.UserType) ([]models.User, error) {
	db := g.db

	var users []models.User
	result := db.Model(&models.User{}).
		Preload(clause.Associations).
		Where(&models.User{Type: usersType}).
		Find(&users)

	return users, result.Error
}

func (g GormStorage) UpdateUserData(user *models.User) error {
	db := g.db

	result := db.Model(&models.User{Model: gorm.Model{ID: user.ID}}).Updates(models.User{
		Model:   gorm.Model{ID: user.ID},
		Email:   user.Email,
		Name:    user.Name,
		Picture: user.Picture,
		Type:    user.Type,
	})

	return result.Error
}
