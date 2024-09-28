package advanced

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func serializeEx() {

	// json.Marshal -> 序列化

	apple := Product{
		Id:       1,
		Name:     "Apple",
		Quantity: 30,
	}

	text, err := json.Marshal(apple)
	if err != nil {
		fmt.Println("JSON serialiation error : ", err)
		return
	}

	fmt.Println(string(text))
}

func deseralizeEx() {
	jsonStr := `{"id":1,"name":"Apple","quantity":30}`

	var product Product

	err := json.Unmarshal([]byte(jsonStr), &product)
	if err != nil {
		fmt.Println("JSON deserialiation error : ", err)
		return
	}

	fmt.Println(product)
}

func JsonSerialize() {
	serializeEx()
	deseralizeEx()
}
