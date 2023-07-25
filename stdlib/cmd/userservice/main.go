package main

import "github.com/tanveerprottoy/starter-go/stdlib/internal/userservice"

func main() {
	a := userservice.NewApp()
	a.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
