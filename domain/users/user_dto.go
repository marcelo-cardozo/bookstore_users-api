package users

import (
	"github.com/marcelo-cardozo/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	email := strings.TrimSpace(strings.ToLower(user.Email))
	if email == "" {
		return errors.NewBadRequestRestErr("user not valid")
	}
	return nil
}
