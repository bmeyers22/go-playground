package main

import (
	"github.com/bmeyers22/go-playground/plugin"
	"fmt"
)

var (
	MainString = "test"
)

func main() {
	fmt.Printf("Main var: %s\n", MainString)
	plugin.PrintTestVar()
}
