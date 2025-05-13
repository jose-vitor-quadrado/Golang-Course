package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		authInfo.username,
		authInfo.password,
	)
}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("=========================================")
}

func main() {
	user1 := authenticationInfo{
		username: "Bing",
		password: "98765",
	}
	user2 := authenticationInfo{
		username: "DDG",
		password: "76921",
	}

	test(user1)
	test(user2)
}
