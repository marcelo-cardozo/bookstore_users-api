package service

import (
	"github.com/marcelo-cardozo/bookstore_users-api/domain/users"
	"github.com/marcelo-cardozo/bookstore_users-api/utils/errors"
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
}

var (
	UsersService usersServiceInterface
)

func init() {
	UsersService = &usersService{}
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId < 0 {
		return nil, errors.NewBadRequestRestErr("user id should be greater than 0")
	}

	user := users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}

	return &user, nil
}
