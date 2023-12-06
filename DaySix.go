package main

import (
	"fmt"
	"math"
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
	//scan lines into data structure
	for lineidx, line := range lines {
		numStr := ""
		regex := regexp.MustCompile(" ?(\\d+)")
		matches := regex.FindAllStringSubmatch(line, -1)
		gameidx := 0
		for _, match := range matches {
			num, _ := strconv.Atoi(match[1])
			if lineidx == 0 {
				game := gameDaysix{num, 0}
				games = append(games, game)
			} else {
				games[gameidx].recordLen = num
				gameidx++
			}
			numStr += match[1]
		}
		num, _ := strconv.Atoi(numStr)
		if lineidx == 0 {
			gamept2.time = num
		} else {
			gamept2.recordLen = num
		}
	}
	//part1
	day.outputpart1 = 1
	comb := 1
	for _, game := range games {
		comb *= combinationsMath(game)
		day.outputpart1 = comb
	}
	//pt2 - bruteforce should probably not be done but it is < 100 ms run time so we keep it simple
	day.outputpart2 = combinationsMath(gamept2)
	
	fmt.Printf("Day six part 1: %d, part 2: %d \n", day.outputpart1, day.outputpart2)
}



func combinationsMath(game gameDaysix) int {
	//EXPLAINER we want to find all games that hold score > recordLen
	//we know output = holdTime * (totalTime - holdTime)
	//thus we want holdTime*(totalTime - holdTime) > recordLen 
	//-holdTime^2+totalTime*holdTime - recordLen > 0 
	//this means we can use the quadratic equation ax^2+bx+c
	discriminantVal := discriminant(game.time, game.recordLen)
	minval := 0
	minfl := 0.0
	maxfl := 0.0
	maxval := 1
	if discriminantVal > 0 {
		//-b +- sqrt(d) / 2 a
		maxfl = (float64(-game.time) - math.Sqrt(float64(discriminantVal))) / (-2.0)
		minfl = (float64(-game.time) + math.Sqrt(float64(discriminantVal))) / (-2.0)

		//say we have 20.00 this means the max value we can pick is actually 19 whereas values like 5.3 has 5 as their max value. credit goes to https://www.reddit.com/user/reddit_Twit/
		//for handling my edge case here.
		maxval = int(math.Ceil(maxfl - 1))
		minval = int(math.Floor(minfl + 1))
	}
	fmt.Printf("min %d max %d\n", minval, maxval)
	return maxval - minval + 1
}

// val - record> -x^2+bx-record
func discriminant(totalTime int, record int) int {
	//discrimant of f(x) = -x^2+bx-c, where holdTime = x and totalTime = b-x
	//a = -1 b = totalTime c = record
	//d = totalTime
	return totalTime*totalTime - 4*(record)
}
