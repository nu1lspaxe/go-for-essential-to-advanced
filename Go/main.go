package main

import (
	"fmt"

	"github.com/nu1lspaxe/go-for-essential-to-advanced/Go/advanced"
)

func main() {
	itOne := advanced.NewItem(advanced.SetId(0), advanced.SetValue("C#"))
	itTwo := advanced.NewItem(advanced.SetId(1))

	fmt.Println(itOne)
	fmt.Println(itTwo)
}
