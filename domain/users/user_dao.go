package users

import (
	"github.com/marcelo-cardozo/bookstore_users-api/utils/date_utils"
	"github.com/marcelo-cardozo/bookstore_users-api/utils/errors"
)

var (
	Db = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	current := Db[user.Id]
	if current == nil {
		return errors.NewNotFoundRestErr("user not in database")
	}
	user.FirstName = current.FirstName
	user.LastName = current.LastName
	user.Email = current.Email
	user.DateCreated = current.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	user.DateCreated = date_utils.GetNowString()

	Db[user.Id] = user
	return nil
}
