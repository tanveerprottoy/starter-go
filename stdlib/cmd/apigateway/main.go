package main

import "github.com/tanveerprottoy/starter-go/stdlib/internal/app/apigateway"

func main() {
	a := apigateway.NewApp()
	a.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
