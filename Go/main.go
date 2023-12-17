package main

import (
	advanced "advanced"
	"fmt"
	"reflect"
)

func main() {
	advanced.RunPrint()

	test := "dsfa"

	if !reflect.ValueOf(test).IsZero() {
		fmt.Println("hi")
	}
}
