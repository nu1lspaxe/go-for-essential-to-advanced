package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	_ "github.com/joho/godotenv"
)

func InitStatusFile() []*os.File {
	err := os.RemoveAll("./status/")
	if err != nil {
		log.Fatalf("error occurs when removing folder: %v\n", err)
	}

	err = os.Mkdir("status", 0750)
	if err != nil {
		log.Fatalf("error occurs when making folder: %v\n", err)
	}

	var statusFiles = make([]*os.File, 4)

	for i := 0; i < 4; i++ {
		filename := fmt.Sprintf("./status/%d.txt", i)
		file := CreateFile(filename)
		statusFiles[i] = file
		file.Close()
	}
	return statusFiles
}

func CreateFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed to create file...")
	}
	return file
}

func WriteIndi(file *os.File, faan []int) {
	file.WriteString("階段: 翻開寶牌指示牌\n")
	file.Write([]byte(fmt.Sprintf("其他玩家是否牌面有番數增加: %v\n", faan)))
}

func OpenStatusFile(index int) *os.File {
	filename := fmt.Sprintf("./status/%d.txt", index)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		log.Fatalf("failed to open status file...")
	}
	return file
}

func WriteA(file *os.File, faan []int) {
	defer file.Close()
	file.WriteString("階段: 翻開寶牌指示牌\n")
	file.Write([]byte(fmt.Sprintf("其他玩家是否牌面有番數增加: %v\n", faan)))
}

func main() {
	TestFn(1, []bool{false, true})
	TestFn(3)
}

func TestFn(test int, tt ...[]bool) {
	if len(tt) > 0 {
		for _, t := range tt[0] {
			if t {
				fmt.Println("t")
			}
		}
		fmt.Println("tt > 0")
	} else {
		fmt.Println("tt <= 0")
	}
}

func GetRanking(scores []int, playerScore int) int {
	copyScores := make([]int, len(scores))
	copy(copyScores, scores)

	sort.SliceStable(copyScores, func(i, j int) bool {
		return copyScores[i] > copyScores[j]
	})

	for i := 0; i < len(scores); i++ {
		if copyScores[i] == playerScore {
			return (i + 1)
		}
	}

	return -1
}
