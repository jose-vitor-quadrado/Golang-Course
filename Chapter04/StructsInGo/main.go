package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("======================================================================")
}

func main() {
	test(messageToSend{
		phoneNumber: 1231982318238,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 7823157581207,
		message:     "Love to have you aboard!",
	})
	test(messageToSend{
		phoneNumber: 1434244897053,
		message:     "We're so excited to have you",
	})
}
