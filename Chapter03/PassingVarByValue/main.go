package main

import "fmt"

func incrementSeeds(sendsSoFar, sendsToAdd int) int {
	sendsSoFar += sendsToAdd
	return sendsSoFar
}

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSeeds(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
}
