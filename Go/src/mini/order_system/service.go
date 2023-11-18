package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字串輸入擷取功能
func scanner(sen string) string {
	var input string
	fmt.Print(sen)
	fmt.Scanln(&input)
	return input
}

// 系統代號一，新增品項
func menuAdd() {

	var drink ITEM
	drink.name = scanner("請輸入要新增的名稱：")

	// 檢查品項是否重複
	for i := 0; i < len(itemList.items); i++ {
		if drink.name == itemList.items[i].name {
			fmt.Printf("執行錯誤。原因：「 %s 」名稱重複\n", drink.name)
			return
		}
	}
	drink.price = editNum("請輸入要新增的價格：")
	itemList.itemAdd(drink)
}

// 系統代號二，檢視品項列表
func menuView(itemlist []ITEM) bool {

	// 顯示清單開頭
	fmt.Println(strings.Repeat("*", 9), " Menu ", strings.Repeat("*", 9))

	// 檢查清單里內是否有品項
	if len(itemlist) == 0 {
		fmt.Println("目前沒有任何飲料項目")
		return false

	} else {
		for i := 0; i < len(itemlist); i++ {
			fmt.Printf("(%d)\t%s ... $ %d\n", i+1, itemlist[i].name, itemlist[i].price)
		}
		return true
	}
}

// 偵錯+調整產品數量、偵錯價格
func editNum(sen string) int {
	for {
		num, _ := strconv.Atoi(scanner(sen))
		if num > 0 {
			return num
		}
		fmt.Println("執行錯誤，請再輸入一次。原因：不得小於零(非正整數)")
	}
}

// 偵錯功能(用於品項、冰量、甜度)
func ifOver(sen string, num int) int {
	for {
		code, _ := strconv.Atoi(scanner(sen))
		if 0 < code && code <= num {
			return code
		}
		fmt.Println("執行錯誤，請再輸入一次。原因：無此選項")
	}
}

// 調整冰量功能
func editIce() string {

	iceList := map[int]string{1: "正常", 2: "少冰", 3: "去冰"}
	fmt.Println("請選擇冰量")

	for i := 0; i < len(iceList); i++ {
		fmt.Printf("(%d)\t%s\n", i+1, iceList[i+1])
	}

	return iceList[ifOver("請輸入編號：", len(iceList))]
}

// 調整甜度功能
func editSugar() string {

	sugarList := map[int]string{1: "正常", 2: "少糖", 3: "無糖"}
	fmt.Println("請選擇甜度")

	for i := 0; i < len(sugarList); i++ {
		fmt.Printf("(%d)\t%s\n", i+1, sugarList[i+1])
	}

	return sugarList[ifOver("請輸入編號：", len(sugarList))]
}

// 檢查購買品項是否重複
func ifSame(item ITEM) bool {

	for i := 0; i < len(cusList.items); i++ {

		if (item.name == cusList.items[i].name) && (item.ice == cusList.items[i].ice) && (item.sugar == cusList.items[i].sugar) {

			cusList.items[i].number = cusList.items[i].number + item.number
			fmt.Println("已有重複品項，已將數量進行加總")
			return true
		}
	}
	return false
}

// 系統代號三，品項銷售
func menuSell() {

	if menuView(itemList.items) {

		var drinkBuy ITEM

		// 購買之品項
		drinkCode := ifOver("請輸入編號：", len(itemList.items)) - 1
		drinkBuy.name = itemList.items[drinkCode].name
		drinkBuy.price = itemList.items[drinkCode].price

		// 購買之數量
		drinkBuy.number = editNum("請輸入數量：")

		// 購買之冰量
		drinkBuy.ice = editIce()

		// 購買之甜度
		drinkBuy.sugar = editSugar()

		// 將品項添加到購物車中
		if !ifSame(drinkBuy) {
			cusList.itemAdd(drinkBuy)
			fmt.Println("購買成功")
		}
	}
}

// 系統代號四，檢視當前訂單內容
func cartView(cuslist []ITEM) bool {
	if len(cuslist) == 0 {
		fmt.Println("購物車內無品項")
		return false
	} else {
		fmt.Printf("%s\t%-4s\t\t%s\t%s\t%s\n", "編號", "品名", "數量", "甜度", "冰量")
		for i := 0; i < len(cuslist); i++ {
			fmt.Printf("%04d\t%-6s\t%d\t%s\t%s\n", i+1, cuslist[i].name, cuslist[i].number, cuslist[i].sugar, cuslist[i].ice)
		}
		return true
	}
}

// 系統代號五，編輯訂單內容
func cartEdit() {
	if cartView(cusList.items) {

		code := ifOver("請輸入編號進行編輯：", len(cusList.items))
		editItem := cusList.items[code-1]
		// 將編輯項目暫時移出購物清單
		cusList.itemRemove(code - 1)

		fmt.Println("(1)	編輯冰量\n" +
			"(2)	編輯甜度\n" +
			"(3)	編輯數量\n" +
			"(4)	刪除品項\n" +
			"(0)	上一頁")

		edit := scanner("請輸入號碼選擇項目：")

		if edit == "1" {
			// 編輯冰量
			editItem.ice = editIce()

		} else if edit == "2" {
			// 編輯甜度
			editItem.sugar = editSugar()

		} else if edit == "3" {
			// 編輯數量
			editItem.number = editNum("請輸入數量：")

		} else if edit == "4" {
			// 刪除品項
			input := scanner("是否刪除(Y/N):")

			if input == "Y" || input == "y" {
				fmt.Println("刪除成功")
			} else {
				fmt.Println("刪除失敗")
				cusList.itemInsert(editItem, code-1)
			}
			return

		} else if edit == "0" {
			// 返回上頁
			cusList.itemInsert(editItem, code-1)
			cartEdit()
			return

		} else {
			fmt.Println("輸入非在範圍內號碼")
			cusList.itemInsert(editItem, code-1)
			return
		}

		if (edit == "1" && !ifSame(editItem)) || (edit == "2" && !ifSame(editItem)) || (edit == "3") {
			fmt.Println("修改成功")
			cusList.itemInsert(editItem, code-1)
		}
	}
}

// 系統代號零，總金額計算
func amount(list ITEMLIST) int {

	total := 0

	for _, it := range list.items {
		total = total + it.number*it.price
	}

	return total
}
