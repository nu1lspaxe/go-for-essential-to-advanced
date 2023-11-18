package advanced

import (
	"fmt"
)

func buy(customer string, products ...string) {
	if products == nil {
		fmt.Printf("%s didn't buy any thing.\n", customer)
		return
	}

	for _, product := range products {
		fmt.Printf("%s buy %s.\n", customer, product)
	}
}

func ChangeableVar() {
	// [...]int -> 可變長度; index:value -> 指定 value 位於 index
	lis := [...]int{1: 10, 9: 5}

	fmt.Println(lis)

	buy("John", "Apple", "Orange")

}
