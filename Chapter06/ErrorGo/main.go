package main

import "fmt"

func getUser() (string, error) {
	return "Dongalo", nil
}

func main() {
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}
