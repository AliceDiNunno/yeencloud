package postgres

import (
	"back/src/core/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID string `gorm:"primary_key"`

	Email    string `gorm:"unique"`
	Password string
}

func (db *Database) CountUsers() int64 {
	var count int64
	db.engine.Model(&User{}).Count(&count)
	return count
}

func (db *Database) FindUserByID(id domain.UserID) (domain.User, error) {
	var user User
	result := db.engine.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return userToDomain(user), nil
}

func (db *Database) FindUserByEmail(email string) (domain.User, error) {
	var user User
	result := db.engine.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return userToDomain(user), nil
}

func (db *Database) CreateUser(user domain.User) (domain.User, error) {
	userToCreate := domainToUser(user)

	result := db.engine.Create(&userToCreate)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return userToDomain(userToCreate), nil
}

func userToDomain(user User) domain.User {
	return domain.User{
		ID:       domain.UserID(user.ID),
		Email:    user.Email,
		Password: user.Password,
	}
}

func domainToUser(user domain.User) User {
	return User{
		ID:       user.ID.String(),
		Email:    user.Email,
		Password: user.Password,
	}
}
