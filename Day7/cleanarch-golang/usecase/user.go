package usecase

import (
	"cleanarch-golang/entity"
	"cleanarch-golang/repository"
)

type IUserUsecase interface {
	CreateUser(user entity.CreateUserRequest) (entity.User, error)
}

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

func (usecase UserUsecase) CreateUser(user entity.CreateUserRequest) (entity.UserResponse, error) {
	u := entity.User{
		Firstname: user.Firstname,
		Lastname: user.Lastname,
		Email: user.Email,
	}

	users, err := usecase.userRepository.Store(u)

	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		Id: users.Id,
		Firstname: users.Firstname,
		Lastname: users.Lastname,
		Email: users.Email,
	}

	return userRes, nil
}