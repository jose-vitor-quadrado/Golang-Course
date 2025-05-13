package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("==============================================")
}

func main() {
	var sender sender
	sender.name = "Sally"
	sender.number = 1231223123
	sender.rateLimit = 1000

	test(sender)
}
