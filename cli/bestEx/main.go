package main

import (
	"fmt"

	"github.com/hangcodebug/golunzi/cli/bestEx/appoptions"
)


func main() {
	NewApp("test").Run()

	fmt.Println(appoptions.NewOptions().MySQLOptions)
}