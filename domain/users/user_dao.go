package users

import (
	"fmt"
	"github.com/bm1905/bookstore_users_api/utils/errors"
)

var usersDB = make(map[int64]*User)

func (user *User) Get() *errors.RestError {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError {
	current := usersDB[user.Id]

	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s aleready exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d aleready exists", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
