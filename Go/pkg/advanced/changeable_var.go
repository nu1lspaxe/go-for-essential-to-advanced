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

// Builder
type Item struct {
	Id    int
	Value string
}

type ItemDef func(*Item)

func SetId(id int) ItemDef {
	return func(it *Item) {
		it.Id = id
	}
}

func SetValue(value string) ItemDef {
	return func(it *Item) {
		it.Value = value
	}
}

func NewItem(options ...ItemDef) *Item {
	item := new(Item)
	for _, opt := range options {
		opt(item)
	}
	return item
}
