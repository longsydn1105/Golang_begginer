package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDay  string
	createdAt time.Time
}

func main() {
	userfirstName := getUserData("Enter your first name: ")
	userlastName := getUserData("Enter your last name: ")
	userbirthday := getUserData("Ennter your birthday: ")

	var appUser User

	appUser = User{
		userfirstName,
		userlastName,
		userbirthday,
		time.Now(),
	}

	outputUserDetails(appUser)
}

func getUserData(param string) string {
	fmt.Print(param)
	var userData string
	fmt.Scan(&userData)
	return userData
}

func outputUserDetails(user User) {
	fmt.Println(user.lastName, user.firstName, user.birthDay, user.createdAt)
}
