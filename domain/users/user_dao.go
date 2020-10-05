package users

import (
	"github.com/marcelo-cardozo/bookstore_users-api/datasources/mysql/users_db"
	"github.com/marcelo-cardozo/bookstore_users-api/utils/date_utils"
	"github.com/marcelo-cardozo/bookstore_users-api/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "users_email_uindex"
	queryInsertUser = "insert into users(first_name, last_name, email, date_created) values (?, ?, ?, ?)"
)


func (user *User) Get() *errors.RestErr {
	//current := Db[user.Id]
	//if current == nil {
	//	return errors.NewNotFoundRestErr("user not in database")
	//}
	//user.FirstName = current.FirstName
	//user.LastName = current.LastName
	//user.Email = current.Email
	//user.DateCreated = current.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerErrRestErr(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestRestErr("Email already used")
		}
		return errors.NewInternalServerErrRestErr(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerErrRestErr("error while saving user")
	}
	user.Id = id
	return nil
}
