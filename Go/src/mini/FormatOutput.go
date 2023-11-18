package mini

import (
	"fmt"
	"strings"
)

// 孩童票function
func children(fee int) int {
	return fee - 40
}

// 學生票function
func students(fee int) float64 {
	return float64(fee) * 0.8
}

func FormatOutput() {

	// 命名變數: 站票 = stand_up, 坐票 = sit_down
	var stand_up, sit_down int

	// 變數賦值
	stand_up = 100
	sit_down = 149

	// 標題列印
	fmt.Println(strings.Repeat("*", 20), "火車售票亭", strings.Repeat("*", 20))
	fmt.Printf("%-2s\t%2s\t%-4s\t%-4s\t%-4s\n", "編號", "票種", "全票", "孩童票", "學生票")

	// 項目列印
	fmt.Printf("%02d\t%-2s\t$%4d\t$%4d\t$%4.0f\n", 1, "站票", stand_up, children(stand_up), students(stand_up))
	fmt.Printf("%02d\t%-2s\t$%4d\t$%4d\t$%4.0f\n", 2, "坐票", sit_down, children(sit_down), students(sit_down))

	// 結尾列印
	fmt.Println(strings.Repeat("*", 52))
}
