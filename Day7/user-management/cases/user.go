package cases

import (
	"user-management/entity"
	"user-management/repository"
)

type IUserCases interface {
	CreateUser(user entity.CreateUserRequest) (entity.User, error)
	GetAllUser() ([]entity.User, error)
	GetUserById(id int) (entity.User, error)
	// UpdateUser(userRequest entity.UpdateUserRequest, id int) (entity.User, error)
	// DeleteUser(id int) error
}

type UserCases struct {
	userRepository repository.IUserRepository
}

func NewUserCases(userRepository repository.IUserRepository) *UserCases {
	return &UserCases{userRepository}
}

func (cases UserCases) CreateUser(user entity.CreateUserRequest) (entity.UserResponse, error) {
	u := entity.User{
		Firstname: user.Firstname,
		Lastname: user.Lastname,
		Email: user.Email,
		Gender: user.Gender,
	}

	users, err := cases.userRepository.Store(u)

	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		Id: users.Id,
		Firstname: users.Firstname,
		Lastname: users.Lastname,
		Email: users.Email,
		Gender: users.Gender,
	}

	return userRes, nil
}

func (cases UserCases) GetAllUser() ([]entity.UserResponse, error) {
	users, err := cases.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	userRes := []entity.UserResponse{}
	for _, user := range users{
		appendUser := entity.UserResponse{
			Id: user.Id,
			Firstname: user.Firstname,
			Lastname: user.Lastname,
			Email: user.Email,
			Gender: user.Gender,
		}
		userRes = append(userRes, appendUser)
	}

	return userRes, nil
}

func (cases UserCases) GetUserById(id int) (entity.UserResponse, error) {
	user, err := cases.userRepository.FindById(id)
	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		Id: user.Id,
		Firstname: user.Firstname,
		Lastname: user.Lastname,
		Email: user.Email,
		Gender: user.Gender,
	}

	return userRes, nil
}