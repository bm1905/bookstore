package users

import (
	"strings"

	"github.com/bm1905/bookstore_users_api/utils/errors_utils"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

const (
	StatusActive = "active"
)

func (user *User) Validate() *errors_utils.RestError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	if user.FirstName == "" {
		return errors_utils.NewBadRequestError("Invalid FirstName")
	}
	user.LastName = strings.TrimSpace(user.LastName)
	if user.LastName == "" {
		return errors_utils.NewBadRequestError("Invalid LastName")
	}
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors_utils.NewBadRequestError("Invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors_utils.NewBadRequestError("Invalid password")
	}
	return nil
}
