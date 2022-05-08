package services

import (
	"github.com/bm1905/bookstore_users_api/domain/users"
	"github.com/bm1905/bookstore_users_api/utils/errors_utils"
)

func CreateUser(user users.User) (*users.User, *errors_utils.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(userId int64) *errors_utils.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors_utils.RestError) {
	current, err := GetUser(user.Id)
	if err != nil {
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

func GetUser(userId int64) (*users.User, *errors_utils.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func SearchUser() (*users.User, error) {
	return nil, nil
}
