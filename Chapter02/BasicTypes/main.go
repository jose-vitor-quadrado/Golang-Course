package main

import "fmt"

/*
	bool

	string

	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte  // alias for uint8

	rune  // alias for int32
		  // represents a Unicode code point

	float32 float64

	complex64 complex128
*/

func main() {
	var smsSendingLimit int = 300
	var costPerSMS float64 = 0.05
	var hasPermission bool = true
	var username string = "sherb"

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
	)
}
