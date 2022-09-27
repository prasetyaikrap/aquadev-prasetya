package repository

import (
	"user-management/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Store(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindById(id int) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(id int) error
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

func (repo UserRepository) FindAll() ([]entity.User, error) {
	users := []entity.User{}
	if err := repo.db.Debug().Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo UserRepository) FindById(id int) (entity.User, error) {
	var user entity.User
	if err := repo.db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (repo UserRepository) Update(user entity.User) (entity.User, error) {
	if err := repo.db.Debug().Save(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (repo UserRepository) Delete(id int) error {
	if err := repo.db.Debug().Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}