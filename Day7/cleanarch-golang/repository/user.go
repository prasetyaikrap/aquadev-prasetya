package repository

import (
	"cleanarch-golang/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Store(user entity.User) (entity.User, error)
	// FindAll() ([]entity.User, error)
	// FindById(id int) (entity.User, error)
	// Update(user entity.User) (entity.User, error)
	// Delete(id int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo UserRepository) Store(user entity.User) (entity.User, error) {
	if err := repo.db.Debug().Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}