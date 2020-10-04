package users

import (
	"github.com/marcelo-cardozo/bookstore_users-api/utils/errors"
	"time"
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
	user.DateCreated = time.Now().String()

	Db[user.Id] = user
	return nil
}
