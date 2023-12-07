package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	dayArray := []Day{
		DayOne{inputPath: "./inputs/dayOne.txt"},
		DayTwo{inputPath: "./inputs/dayTwo.txt"},
		DayThree{inputPath: "./inputs/dayThree.txt"},
		DayFour{inputPath: "./inputs/DayFour.txt"},
		DayFive{inputPath: "./inputs/DayFive.txt"},
		DaySix{inputPath: "./inputs/DaySix.txt"},
		DaySeven{inputPath: "./inputs/DaySeven.txt"},
		DayEight{},
		DayNine{},
		DayTen{},
		DayEleven{},
		DayTwelve{},
		DayThirteen{},
		DayFourteen{},
		DayFifteen{},
		DaySixteen{},
		DaySeventeen{},
		DayEighteen{},
		DayNineteen{},
		DayTwenty{},
		DayTwentyOne{},
		DayTwentyTwo{},
		DayTwentyThree{},
		DayTwentyFour{},
		DayTwentyFive{},
		//new days here
	}

	onlyToday := false
	includeTo := 10000
	includeFrom := 0
	trackExecutionTime := false
	if len(os.Args[1:]) > 0 {
		for _, arg := range os.Args {
			if arg == "-today" {
				onlyToday = true
			} else {
				regStr := "-includeFrom=(\\d*)"
				regex := regexp.MustCompile(regStr)
				res, _ := regexp.MatchString(regStr, arg)
				if res {
					match := regex.FindAllSubmatch([]byte(os.Args[1]), -1)
					num, _ := strconv.Atoi(string(match[0][1]))
					includeFrom = num
				} else {
					regStr := "-includeTo=(\\d*)"
					regex := regexp.MustCompile(regStr)
					res, _ := regexp.MatchString(regStr, arg)
					if res {
						match := regex.FindAllSubmatch([]byte(arg), -1)
						num, _ := strconv.Atoi(string(match[0][1]))
						includeTo = num
					} else {
						regStr := "-includeTime"
						res, _ := regexp.MatchString(regStr, arg)
						if res {
							trackExecutionTime = true
						}
					}
				}
			}

		}
	}
	var startTime time.Time
	if trackExecutionTime {
		startTime = time.Now()
	}

	if onlyToday {
		timeObj := time.Now()
		currDay := timeObj.Day()
		if currDay > 25 {
			println("sorry today functionality is only available from 1-25th december")
		} else {
			dayArray[currDay-1].Solve()
		}
	} else {
		for index, day := range dayArray {
			if index < includeFrom-1 {
				continue
			}
			if index > includeTo-1 {
				continue
			}
			day.Solve()
		}
	}
	if trackExecutionTime {
		endTime := time.Now()
		execTime := endTime.Sub(startTime)
		fmt.Printf("total execution time: %v \n", execTime)
	}
}
