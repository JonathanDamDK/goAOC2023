package main

import (
	"fmt"
	"regexp"
	"strconv"
)
type DaySix struct {
	inputPath string
	input     []string
	outputpart1    int
	
}
type gameDaysix struct{
	time  int
	recordLen int
}
func (day DaySix) Solve() {
	lines := MapFileToStringArr(day.inputPath)
	var games  []gameDaysix
	for lineidx, line := range lines {
		regex := regexp.MustCompile(" ?(\\d+)")
		matches := regex.FindAllStringSubmatch(line, -1)
		for matchidx, match := range matches{
			fmt.Printf("%v \n",match)
			num, _ := strconv.Atoi(match[1])
			if(lineidx  == 0){
				game := gameDaysix{num,0}
				games = append(games,game)
			} else {
				games[matchidx].recordLen = num
			}
		}
	}
	fmt.Printf("%v \n", games)
	fmt.Printf("Day six part 1: %d \n", day.outputpart1)
}
