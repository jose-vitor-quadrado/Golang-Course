package main 

import "fmt"

myCar := struct {
	Make string
	Model string
} {
	Make: "tesla",
	Model: "model 3"
}

func main() {
	fmt.PrintLn(myCar.Model)
}