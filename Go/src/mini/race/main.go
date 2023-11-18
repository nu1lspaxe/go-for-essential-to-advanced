package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

type Competition interface {
	Running()
}

type Competitor struct {
	Name     string
	Speed    int
	Time     int
	Runned   int
	RestDis  int
	RestTime int
	LeftTime int
}

var wg sync.WaitGroup
var distance int = 100

func ReadFile(filename string) []Competitor {
	file, _ := os.ReadFile(filename)
	lines := strings.Split(string(file), "\n")
	var competitors []Competitor
	for _, line := range lines {
		var competitor Competitor
		fmt.Sscan(strings.Replace(line, ",", " ", -1), &competitor.Name, &competitor.Speed, &competitor.RestDis, &competitor.RestTime)
		competitors = append(competitors, competitor)
	}
	return competitors
}

func (c *Competitor) Running() {

	defer wg.Done()

	if c.Runned == distance {
		return
	}
	c.Time++
	if c.LeftTime > 0 {
		c.LeftTime--
		return
	}
	c.Runned += c.Speed
	if (c.RestTime > 0) && (c.Runned%c.RestDis == 0) {
		c.LeftTime = c.RestTime
	}
	if c.Runned > distance {
		c.Runned = distance
	}
}

func NowRunned(competitors []Competitor) {
	fmt.Printf("名字\t目前位置\n")
	for _, competitor := range competitors {
		fmt.Printf("%s\t%d公尺\n", competitor.Name, competitor.Runned)
	}
}

func main() {

	competitors := ReadFile("./list.txt")
	fmt.Println("選手就位")
	NowRunned(competitors)

	for {
		var input string
		fmt.Println("請輸入1開始：")
		fmt.Scanln(&input)
		if input == "1" {
			fmt.Println("比賽開始")
			break
		} else {
			fmt.Println("輸入錯誤")
		}
	}

	times := 0
	// fmt.Println(competitors)

	for {
		keep := func() bool {
			for _, competitor := range competitors {
				if competitor.Runned < distance {
					return true
				}
			}
			return false
		}

		if keep() {
			times++
			fmt.Printf("第%d秒\n", times)
			for index, _ := range competitors {

				/* !!!!! 為什麼不能用 _, competitor (struct 會無法儲存計算) !!!!! */
				var c Competition = &competitors[index]
				// go competitors[index].Running()
				wg.Add(1)
				go c.Running()
				// fmt.Println(competitor)
			}
			wg.Wait()
			NowRunned(competitors)
		} else {
			fmt.Println("比賽結束，成績結算")
			fmt.Printf("名次\t名字\t完成秒數\n")
			sort.SliceStable(competitors, func(x, y int) bool {
				return competitors[x].Time < competitors[y].Time
			})
			for index, competitor := range competitors {
				fmt.Printf("第%d名\t%s\t%d\n", index+1, competitor.Name, competitor.Time)
			}
			break
		}
	}
}
