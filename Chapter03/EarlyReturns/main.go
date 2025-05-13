package main

func divide(dividend, divisor int) (int, string) {
	if divisor == 0 {
		return 0, "Can't divide by zero"
	}
	return dividend / divisor, "ok"
}

func main() {

}
