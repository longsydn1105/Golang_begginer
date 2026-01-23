package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userfirstName := getUserData("Enter your first name: ")
	userlastName := getUserData("Enter your last name: ")
	userbirthday := getUserData("Ennter your birthday: ")

	var appUser *user.User

	appUser, err := user.New(userfirstName, userlastName, userbirthday)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("longdaica", "12313")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(param string) string {
	fmt.Print(param)
	var userData string
	fmt.Scanln(&userData)
	return userData
}
