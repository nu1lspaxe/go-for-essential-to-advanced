package mini

import "fmt"

func ScoreShow() {

	var name, grade string
	var math, english int
	var average float64

	fmt.Print("請輸入姓名: ")
	fmt.Scanln(&name)

	fmt.Print("請輸入數學成績: ")
	fmt.Scanln(&math)

	fmt.Print("請輸入英文成績: ")
	fmt.Scanln(&english)

	if math > 100 {
		fmt.Println("輸入成績錯誤，程式結束")
	} else if math < 0 {
		fmt.Println("輸入成績錯誤，程式結束")
	} else if english > 100 {
		fmt.Println("輸入成績錯誤，程式結束")
	} else if english < 0 {
		fmt.Println("輸入成績錯誤，程式結束")
	} else {
		average = (float64(math) + float64(english)) / 2

		switch {
		case average >= 80:
			grade = "A"
		case average >= 70:
			grade = "B"
		case average >= 60:
			grade = "C"
		case 60 > average:
			grade = "F"
		}

		fmt.Printf("%-5s\t%-5s\t%-5s\t%-5s\t%-5s\n", "姓名", "數學", "英文", "平均", "等第")
		fmt.Printf("%-s\t%-5d\t%-5d\t%-0.2f\t%-5s\n", name, math, english, average, grade)
	}
}
