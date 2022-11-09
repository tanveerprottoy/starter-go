package main

import "txp/restapistarter/internal/app"

func main() {
	a := new(app.App)
	a.InitComponents()
	a.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
