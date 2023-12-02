package main

import (
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {

	dayArray := []Day{
		DayOne{inputPath: "./inputs/dayOne.txt"},
		DayTwo{inputPath: "./inputs/dayTwo.txt"},
		DayThree{inputPath:  "./dayThree.txt"},
		//new days here
	}

	onlyToday := false
	includeTo := 10000
	includeFrom := 0
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
					}
				}
			}

		}
		if onlyToday{
			timeObj := time.Now()
			currDay := timeObj.Day()
			if(currDay > 25){
				println("sorry today is not a advent of code day")
			} else{ 
				dayArray[currDay-1].Solve()
			}
		} else {
			for index, day := range dayArray {
				if index < includeFrom-1 {
					continue
				}
				if index > includeTo-1{
					continue
				}
				day.Solve()
			}
		}

	} else {
		for _, day := range dayArray {
			day.Solve()
		}
	}
}
