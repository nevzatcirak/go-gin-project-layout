package repository

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nevzatcirak/go-gin-poc/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		connection: db,
	}
}

func (db *userRepository) Save(user domain.User) {
	db.connection.Create(&user)
}

func (db *userRepository) Update(user domain.User) {
	db.connection.Save(&user)
}

func (db *userRepository) Delete(user domain.User) {
	db.connection.Delete(&user)
}

func (db *userRepository) Find(id uint64) domain.User {
	var user domain.User
	user.ID = id
	db.connection.Preload(clause.Associations).Find(&user)
	return user
}

func (db *userRepository) FindAll() []domain.User {
	var users []domain.User
	db.connection.Preload(clause.Associations).Find(&users)
	return users
}
