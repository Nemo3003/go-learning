package main

import "fmt"

func getUser(){
	fmt.Println("I return a user")
	return 
}

// GO allows you to manage each possible error differently without having to nest try...catch blocks
func err(){
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return err
	}

	profile, err := getProfile(user.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
}