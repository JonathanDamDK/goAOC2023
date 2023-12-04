package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type DayFour struct {
	inputPath string
	input     []string
	Outputpt1 int
	Outputpt2 int
}

func (day DayFour) Solve() {
	day.Outputpt1 = 0
	lines := MapFileToStringArr(day.inputPath)
	countArr := make([]int, len(lines))
	//initialize all counts to 1
	for lineIndex  := range lines {
		countArr[lineIndex] = 1

	}
	for lineIndex, line := range lines {
		var winningNumbers, actualNumbers []int
		//split so we can match on winning and on actual
		stringArr := strings.Split(line, "|")
		regex := regexp.MustCompile(" +(\\d+)")
		match := regex.FindAllStringSubmatch(stringArr[0], -1)
		for index, str := range match {
			if index == 0 {
				continue
			} else {
				num, _ := strconv.Atoi(str[1])
				winningNumbers = append(winningNumbers, num)
			}
		}
		//match actual numbers
		match = regex.FindAllStringSubmatch(stringArr[1], -1)
		for _, str := range match {
			num, _ := strconv.Atoi(str[1])
			actualNumbers = append(actualNumbers, num)

		}
		val := 0
		day.Outputpt2 = 0
		//keep count of the matches and dobule the value for part 1
		count := 0
		for _, num := range actualNumbers {
			for _, winningNum := range winningNumbers {
				if num == winningNum {
					if val == 0 {
						count++
						val = 1
					} else {
						count++
						val *= 2
					}
				}
			}
		}
		if count > 0 {
			for i := 1; i <= count; i++ {
				countArr[lineIndex+i] += countArr[lineIndex]
			}
		}

		day.Outputpt1 += val
	}
	day.Outputpt2 = Sum(countArr)
	fmt.Printf("day 4 part 1: %d part 2: %d \n", day.Outputpt1, day.Outputpt2)
}

func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
