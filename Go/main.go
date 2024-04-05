package main

import (
	// concurrency "concurrency"
	advanced "advanced"
	"fmt"
)

func main() {
	itOne := advanced.NewItem(advanced.SetId(0), advanced.SetValue("C#"))
	itTwo := advanced.NewItem(advanced.SetId(1))

	fmt.Println(itOne)
	fmt.Println(itTwo)
}
