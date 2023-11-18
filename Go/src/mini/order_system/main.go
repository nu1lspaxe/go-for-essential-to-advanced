package main

import (
	"fmt"
	"os"
	"strings"
)

var itemList ITEMLIST
var cusList ITEMLIST

func main() {

	// 基礎設定，預設系統連續動作，直到使用者輸入 (0)退出
	ok := true

	// 讀取預設品項列表
	file, _ := os.ReadFile("./menu.txt")
	fileSplit := strings.Split(string(file), "\n")
	for _, i := range fileSplit {
		var item ITEM
		fmt.Sscan(strings.Replace(i, ",", " ", 1), &item.name, &item.price)
		itemList.itemAdd(item)
	}

	for ok {
		// 顯示店家POS系統
		fmt.Println(strings.Repeat("-", 11), "不要對他傻笑 POS系統", strings.Repeat("-", 11))
		fmt.Println("(1)	管理 - 新增飲料品項\n" +
			"(2)	管理 - 檢視菜單\n" +
			"(3)	銷售\n" +
			"(4)	當前訂單內容\n" +
			"(5)	編輯訂單內容\n" +
			"(0)	退出")

		// 使用者輸入系統代號
		systemCode := scanner("請輸入代號：")

		// 判斷使用者輸入之系統代號，並進入代號對應的功能
		if systemCode == "1" {
			menuAdd()

		} else if systemCode == "2" {
			menuView(itemList.items)

		} else if systemCode == "3" {
			menuSell()

		} else if systemCode == "4" {
			cartView(cusList.items)

		} else if systemCode == "5" {
			cartEdit()

		} else if systemCode == "0" {
			ok = false
			fmt.Printf("總金額為%d元\n", amount(cusList))
			fmt.Println("感謝使用")

		} else {
			fmt.Println("執行錯誤。原因：無此選項")
		}
	}
}
