package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type DaySix struct {
	inputPath   string
	input       []string
	outputpart1 int	
	outputpart2 int
}
type gameDaysix struct {
	time      int
	recordLen int
}

func (day DaySix) Solve() {
	lines := MapFileToStringArr(day.inputPath)
	var games []gameDaysix
	var gamept2 gameDaysix
	for lineidx, line := range lines {
		numStr := ""
		regex := regexp.MustCompile(" ?(\\d+)")
		matches := regex.FindAllStringSubmatch(line, -1)
		for matchidx, match := range matches {
			num, _ := strconv.Atoi(match[1])
			if lineidx == 0 {
				game := gameDaysix{num, 0}
				games = append(games, game)
			} else {
				games[matchidx].recordLen = num
			}
			numStr += match[1]
		}
		num, _ := strconv.Atoi(numStr)
		if lineidx == 0 {
			gamept2.time = num
		} else {
			gamept2.recordLen = num
		}
		fmt.Printf("%v \n", gamept2)

		day.outputpart1 = 1
		for _, game := range games {
			uniqueWays := 0
			for num := 0; num <= game.time; num++ {
				speed := num
				timeRemaining := game.time - num
				if speed*timeRemaining > game.recordLen {
					uniqueWays += 1
				}
			}
			day.outputpart1 = day.outputpart1 * uniqueWays
		}
	}

	//pt2

	for num := 0; num <= gamept2.time; num++ {
		speed := num
		timeRemaining := gamept2.time - num
		if speed*timeRemaining > gamept2.recordLen {
			day.outputpart2 += 1
		}
	}

	fmt.Printf("Day six part 1: %d part 2: %d \n", day.outputpart1,day.outputpart2)
}
