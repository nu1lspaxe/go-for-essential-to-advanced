//go:build exclude

package mini

import (
	"fmt"
	"strings"
)

var menu = make(map[string]int)
var keys = make(map[int]string)

func DrinkOrderSystem() {

	for {
		var userInput int

		fmt.Println(strings.Repeat("-", 11), "不要對他傻笑 POS系統", strings.Repeat("-", 11))
		fmt.Println("(1)	管理 - 新增飲料品項")
		fmt.Println("(2)	管理 - 檢視菜單")
		fmt.Println("(3)	銷售")
		fmt.Println("(4)	退出")
		fmt.Print("請輸入代號 : ")
		fmt.Scan(&userInput)
		fmt.Println()

		if userInput == 1 {
			menuNew()
		} else if userInput == 2 {
			menuView()
		} else if userInput == 3 {
			menuSale()
		} else if userInput == 4 {
			fmt.Println("感謝使用")
			break
		} else {
			fmt.Println("執行錯誤。原因：無此選項")
		}
	}
}

func menuNew() {
	var productName string
	var productPrice int

	fmt.Print("請輸入要新增的名稱：")
	fmt.Scan(&productName)

	_, price := menu[productName]
	if true && price {
		fmt.Printf("執行錯誤。原因：「 %s 」名稱重複\n", productName)
		return
	}

	fmt.Print("請輸入要新增的價格：")
	fmt.Scan(&productPrice)

	menu[productName] = productPrice
}

func menuView() bool {
	fmt.Println(strings.Repeat("*", 9), " Menu ", strings.Repeat("*", 9))

	if len(menu) == 0 {
		fmt.Println("目前沒有任何飲料項目")
		return false
	} else {
		index := 1
		for key, _ := range menu {
			fmt.Printf("(%d)\t%s ... $ %d\n", index, key, menu[key])
			keys[index] = key
			index++
		}
	}
	return true
}

func menuSale() {
	if menuView() {
		var productCode int

		fmt.Print("請輸入編號：")
		fmt.Scan(&productCode)

		if productCode <= len(menu) {
			fmt.Printf("「%s」購買成功。請收款 %d 元\n", keys[productCode], menu[keys[productCode]])
		}
	}
}
