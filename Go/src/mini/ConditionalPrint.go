package mini

import (
	"fmt"
	"strings"
)

// 票數 function
func sum_ticket() int {
	var ticket int

	fmt.Print("請輸入購買張數:")
	fmt.Scanln(&ticket)

	return ticket
}

func ConditionalPrint() {

	// 命名變數: 站票 = stand_up, 坐票 = sit_down, 編號 = number, 票種 = typed
	var stand_up, sit_down, number, typed int

	// 變數賦值
	stand_up = 100
	sit_down = 149

	// 標題列印
	fmt.Println(strings.Repeat("*", 20), "火車售票亭", strings.Repeat("*", 20))
	fmt.Printf("%-2s\t%-2s\t%-4s\t%-4s\t%-4s\n", "編號", "票種", "全票", "孩童票", "學生票")

	// 項目列印
	fmt.Printf("%02d\t%-2s\t$%4d\t$%4d\t$%4.0f\n", 1, "站票", stand_up, children(stand_up), students(stand_up))
	fmt.Printf("%02d\t%-2s\t$%4d\t$%4d\t$%4.0f\n", 2, "坐票", sit_down, children(sit_down), students(sit_down))

	// 結尾列印
	fmt.Println(strings.Repeat("*", 52))

	// 編號選擇輸入
	fmt.Print("請輸入編號購買:")
	fmt.Scanln(&number)

	// 編號資訊顯示
	if number == 1 {
		fmt.Println("選擇站票")
		fmt.Println(strings.Repeat("*", 9), "票種", strings.Repeat("*", 9))
		fmt.Printf("%-2s\t%-2s\t%-2s\n", "編號", "票種", "價錢")
		fmt.Printf("%02d\t%-4s\t$%4d\n", 1, "全票", stand_up)
		fmt.Printf("%02d\t%-4s\t$%4d\n", 2, "孩童票", children(stand_up))
		fmt.Printf("%02d\t%-4s\t$%4.0f\n", 3, "學生票", students(stand_up))
		fmt.Println(strings.Repeat("*", 24))

		// 票種選擇輸入
		fmt.Print("請輸入編號購買:")
		fmt.Scanln(&typed)

		// 票種資訊顯示
		if typed == 1 {
			fmt.Println("選擇全票")
			fmt.Printf("總金額為 %d元\n", int(sum_ticket()*stand_up))
			fmt.Println("程式結束")
		} else if typed == 2 {
			fmt.Println("選擇孩童票")
			fmt.Printf("總金額為 %d元\n", int(sum_ticket()*children(stand_up)))
			fmt.Println("程式結束")
		} else if typed == 3 {
			fmt.Println("選擇學生票")
			fmt.Printf("總金額為 %d元\n", int(sum_ticket()*int(students(stand_up))))
			fmt.Println("程式結束")
		} else {
			fmt.Println("無此票編號")
			fmt.Println("程式結束")
		}

	} else if number == 2 {
		fmt.Println("選擇坐票")
		fmt.Println(strings.Repeat("*", 9), "票種", strings.Repeat("*", 9))
		fmt.Printf("%-2s\t%-2s\t%-2s\n", "編號", "票種", "價錢")
		fmt.Printf("%02d\t%-4s\t$%4d\n", 1, "全票", sit_down)
		fmt.Printf("%02d\t%-4s\t$%4d\n", 2, "孩童票", children(sit_down))
		fmt.Printf("%02d\t%-4s\t$%4.0f\n", 3, "學生票", students(sit_down))
		fmt.Println(strings.Repeat("*", 24))

		// 票種選擇輸入
		fmt.Print("請輸入編號購買:")
		fmt.Scanln(&typed)

		// 票種資訊顯示
		if typed == 1 {
			fmt.Println("選擇全票")
			fmt.Printf("總金額為 %d元\n", int(sum_ticket()*sit_down))
			fmt.Println("程式結束")
		} else if typed == 2 {
			fmt.Println("選擇孩童票")
			fmt.Printf("總金額為 %d元\n", int(sum_ticket()*children(sit_down)))
			fmt.Println("程式結束")
		} else if typed == 3 {
			fmt.Println("選擇學生票")
			fmt.Printf("總金額為 %d元\n", int(sum_ticket()*int(students(sit_down))))
			fmt.Println("程式結束")
		} else {
			fmt.Println("無此票編號")
			fmt.Println("程式結束")
		}

	} else {
		fmt.Println("查無此票")
		fmt.Println("程式結束")
	}
}
