package main

import "github.com/tanveerprottoy/starter-go/gin/internal/app/gin"

func main() {
	a := gin.NewApp()
	a.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
