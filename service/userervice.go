package service

import (
	"fmt"
	"strings"
	"testmainmod/dao"
	"testmainmod/model"
)

type userService struct {
	db dao.UserDB
}

type UserService interface {
	CreateUser(model.User) error
	GetUserByEmail(email string) (*model.User, error)
}

func NewUserService(dao dao.UserDB) UserService {

	return &userService{db: dao}
}

func (us *userService) CreateUser(user model.User) error {
	return us.db.CreateUser(user)
}

func (us *userService) GetUserByEmail(email string) (*model.User, error) {
	//Validate the email id
	tokens := strings.Split(email, "@")
	if len(tokens) > 2 {
		return nil, fmt.Errorf("invalid email id  : %s", email)
	}
	return us.db.GetUserByEmail(email)
}
