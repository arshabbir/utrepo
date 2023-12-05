package dao

import (
	"fmt"
	"testmainmod/model"
)

type fakerepo struct{}

func NewFakeDB() UserDB {
	return &fakerepo{}
}

func (u *fakerepo) CreateUser(user model.User) error {
	// Implement the creation logic

	return nil
}

func (u *fakerepo) Close() error {
	return nil
}

func (u *fakerepo) GetUserByEmail(email string) (*model.User, error) {

	if email == "admin@example.com" {
		return &model.User{Id: 1, First_name: "Admin"}, nil
	}

	return nil, fmt.Errorf("no rows in result set")

}

func (u *fakerepo) UpdateUser(Id string, user model.User) error {
	return nil
}

func (u *fakerepo) DeleteUser(Id string) error {
	return nil
}
