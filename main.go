package main

import (
	"fmt"
	"mvc/initializer"
)

func init() {
	initializer.LoadEnvariables()
initializer.ConnectionDatabase()
}
func main() {
	fmt.Println("js")
}
