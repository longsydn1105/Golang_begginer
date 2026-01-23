package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDay  string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func New(userfirstName, userlastName, userbirthday string) (*User, error) {
	if userfirstName == "" || userlastName == "" || userbirthday == "" {
		return nil, errors.New("Pls put your info")
	}

	return &User{
		userfirstName,
		userlastName,
		userbirthday,
		time.Now(),
	}, nil
}

func (user User) OutputUserDetails() {
	fmt.Println(user.lastName, user.firstName, user.birthDay, user.createdAt)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			"ADMIN",
			"ADMIN",
			"---",
			time.Now(),
		},
	}
}
