package plugin

import "fmt"

var TestVar = "test"
var testVar = "test"

func PrintTestVar() {
	fmt.Printf("Public var is is: %s\nPRivate var is: %s\n", TestVar, testVar)
}
