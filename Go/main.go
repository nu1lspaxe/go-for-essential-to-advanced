package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)

}

func findSecondMax(arr []int) int {
	// input []int
	// output second maximum integer

	a, b := arr[0], arr[1]
	if arr[0] < arr[1] {
		a = arr[1]
		b = arr[0]
	}

	for _, i := range arr[2:] {
		if i > a { // maximum
			b = a
			a = i
		} else if i > b && i != a {
			b = i
		}
	}
	return b
}
