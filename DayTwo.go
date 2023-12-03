package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type DayTwo struct {
	inputPath string
	input     []string
	Output    int
}

type game struct {
	gameId     int
	CountArray []colorCount
}
type colorCount struct {
	Red   int
	Green int
	Blue  int
}

func (day DayTwo) Solve() {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	lines := MapFileToStringArr(day.inputPath)
	game := game{}
	ptTwoOutput := 0
	for _, line := range lines {
		game = parseGame(line)
		valid := true
		maxCount := colorCount{0, 0, 0}
		for _, count := range game.CountArray {
			if count.Red > maxRed || count.Blue > maxBlue || count.Green > maxGreen {
				valid = false //day one, invalidated since the max amount of cubes has been exceeded
			}
			//day two, find the max amount of every cube that has been taken out at one time
			if count.Red > maxCount.Red {
				maxCount.Red = count.Red
			}
			if count.Green > maxCount.Green {
				maxCount.Green = count.Green
			}
			if count.Blue > maxCount.Blue {
				maxCount.Blue = count.Blue
			}

		}
		ptTwoOutput += maxCount.Red * maxCount.Green * maxCount.Blue
		if valid {
			day.Output += game.gameId
		}
	}
	fmt.Printf("day 2 part 1: %d part2 : %d \n", day.Output, ptTwoOutput)

}

func parseGame(gameStr string) game {
	//define local variables
	result := game{}
	currCount := colorCount{0, 0, 0}
	var countArr []colorCount

	//define and compile regex
	regstringRGB := "(\\d*) (blue|red|green)(,|;)?"
	regstringGame := "Game (\\d*):"
	regExRGB := regexp.MustCompile(regstringRGB)
	regExGame := regexp.MustCompile(regstringGame)

	//match to find the game id
	resultGame := regExGame.FindAllSubmatch([]byte(gameStr), -1)
	num, _ := strconv.Atoi(string(resultGame[0][1]))
	result.gameId = num
	resultRGB := regExRGB.FindAllSubmatch([]byte(gameStr), -1)
	//loop through all regex matches, note here that match[0] always containes the whole string
	for _, match := range resultRGB {
		number, _ := strconv.Atoi(string(match[1]))
		//add the count to the correct color
		switch string(match[2]) {
		case "red":
			currCount.Red += number
			break
		case "green":
			currCount.Green += number
			break
		case "blue":
			currCount.Blue += number
			break
		}
		if string(match[3]) == ";" {
			//new count needs to be initialized so we add the current count to the final count array and start counting a new one
			countArr = append(countArr, currCount)
			currCount = colorCount{0, 0, 0}
		}
	}
	//check if there is a count pending, in the case of ; 2 green, there would be a currCount object with a count greater than one,
	if currCount.Red != 0 || currCount.Green != 0 || currCount.Blue != 0 {
		countArr = append(countArr, currCount)
		currCount = colorCount{0, 0, 0} // emphazising that the currCount is reset, is not necessary for

	}
	result.CountArray = countArr
	return result
}
