package services

import (
	"github.com/bm1905/bookstore_users_api/domain/users"
	"github.com/bm1905/bookstore_users_api/utils/crypto_utils"
	"github.com/bm1905/bookstore_users_api/utils/dates_utils"
	"github.com/bm1905/bookstore_users_api/utils/errors_utils"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors_utils.RestError)
	CreateUser(users.User) (*users.User, *errors_utils.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors_utils.RestError)
	DeleteUser(int64) *errors_utils.RestError
	SearchUser(string) (users.Users, *errors_utils.RestError)
	GetAllUsers() (users.Users, *errors_utils.RestError)
}

func (s *userService) GetUser(userId int64) (*users.User, *errors_utils.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors_utils.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = dates_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors_utils.RestError) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *userService) DeleteUser(userId int64) *errors_utils.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *userService) SearchUser(status string) (users.Users, *errors_utils.RestError) {
	users := &users.User{}
	return users.FindByStatus(status)
}

func (s *userService) GetAllUsers() (users.Users, *errors_utils.RestError) {
	users := &users.User{}
	return users.GetAll()
}
