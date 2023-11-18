package mini

import (
	"fmt"
	"sort"
)

func ScoreRecord() {

	var member int
	var name []string
	var math []int
	var english []int
	var average []float64
	var grade []string

	for {
		fmt.Print("請輸入需要紀錄幾位學生:")
		fmt.Scanln(&member)

		if member > 0 {
			break
		}
		fmt.Println("請輸入正整數")
	}

	for i := 0; i < member; i++ {
		var name_reg string
		var math_reg, english_reg int
		var average_reg float64

		for {
			fmt.Print("請輸入姓名: ")
			fmt.Scanln(&name_reg)

			fmt.Print("請輸入數學成績: ")
			fmt.Scanln(&math_reg)

			fmt.Print("請輸入英文成績: ")
			fmt.Scanln(&english_reg)
			if math_reg <= 100 && math_reg >= 0 && english_reg <= 100 && english_reg >= 0 {
				break
			}
			fmt.Println("輸入成績錯誤，請再輸入一次")
		}

		name = append(name, name_reg)
		math = append(math, math_reg)
		english = append(english, english_reg)
		average_reg = (float64(math_reg) + float64(english_reg)) / 2
		average = append(average, average_reg)

		switch {
		case average_reg >= 80:
			grade = append(grade, "A")
		case average_reg >= 70:
			grade = append(grade, "B")
		case average_reg >= 60:
			grade = append(grade, "C")
		case 60 > average_reg:
			grade = append(grade, "F")
		}
	}

	fmt.Printf("%-5s\t%-5s\t%-5s\t%-5s\t%-5s\n", "姓名", "數學", "英文", "平均", "等第")

	for i := 0; i < member; i++ {
		fmt.Printf("%-s\t%-5d\t%-5d\t%-0.1f\t%-5s\n", name[i], math[i], english[i], average[i], grade[i])
	}

	student := make(map[float64]string)

	for i := 0; i < member; i++ {
		student[average[i]] = name[i]
	}

	fmt.Println("全班成績排名依序為")
	sort.Float64s(average)
	for i := 0; i < member; i++ {
		fmt.Println("第", i+1, "名: ", student[average[len(average)-(i+1)]])
	}

}
