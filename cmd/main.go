package main

import "txp/restapistarter/app"

func main() {
	app := new(app.App)
	app.InitComponents()
	app.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
