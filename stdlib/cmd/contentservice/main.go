package main

import "github.com/tanveerprottoy/starter-go/stdlib/internal/contentservice"

func main() {
	a := contentservice.NewApp()
	a.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
