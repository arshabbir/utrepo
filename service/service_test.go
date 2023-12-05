package service

import (
	"errors"
	"log"
	"testing"
	"testmainmod/dao"
	"testmainmod/model"
)

var us UserService

func TestMain(m *testing.M) {
	// Setup
	log.Println("MainService : STarted ////")
	setup(m)
	// Run tests
	m.Run()
	//Teardown
	teardown(m)
	log.Println("MainService : Ended ////")
}

func TestCreateuser(t *testing.T) {
	log.Println("T testing started ")
	err := us.CreateUser(model.User{Id: 7, First_name: "f1", Last_name: "L1", Email: "arshabbir@gmail.com"})
	if err != nil {
		t.Errorf("Failed ")
	}
}

// TODO : Continue testing with testing.M (setup and tear down and then start the mocking of the db )

func TestGetUserByEmail(t *testing.T) {
	log.Println("T testing started ")
	tests := []struct {
		name       string
		email      string
		expected   *model.User
		err        error
		errMessage string
	}{
		{"valied email id", "admin@example.com", &model.User{First_name: "Admin"}, nil, ""},
		{"email id not exists", "notexits@emaial.com", nil, errors.New(""), "no rows in result set"},
		//{"invalied ", "arshabbir ", nil, errors.New(""), "invalid email id  : arshabbir"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			user, err := us.GetUserByEmail(test.email)
			if test.err == nil {
				if err != nil {
					t.Errorf("Failed %s", err.Error())
				}
				if test.expected.First_name != user.First_name {
					t.Errorf("Unexpected user name received : %s", user.First_name)
				}
			}

			if test.err != nil {
				if test.errMessage != err.Error() {
					t.Errorf("Failed .... %s", err.Error())
				}
			}
		})

	}

}

func setup(m *testing.M) {
	log.Println("****************setup ")
	//	addr := ":8080"

	us = NewUserService(dao.NewFakeDB())
}

func teardown(m *testing.M) {
	log.Println("****************Teardown")
}

func FuzzGetUserByemailid(f *testing.F) {
	f.Add("arshabbir@gmail.com")
	f.Fuzz(func(t *testing.T, id string) {
		us.GetUserByEmail(id)
	})
}
