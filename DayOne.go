package main

import (
	"fmt"
	"regexp"
)

type DayOne struct {
	inputPath  string
	input      []string
	output     int
	output_snd int
}

func (day DayOne) Solve() {

	//part one
	day.output = 0
	day.input = MapFileToStringArr(day.inputPath)
	for _, elem := range day.input {

		///regex meaning two groups of 1 digit, sorrounded by any characters repeated 0 or more times
		reg := regexp.MustCompile("([1-9]).*([1-9]).*")
		match := reg.FindStringSubmatch(elem)
		if len(match) > 0 {
			num1 := MapStringToInt(match[1])
			num2 := MapStringToInt(match[2])
			day.output += num1*10 + num2
		} else {
			reg := regexp.MustCompile("([1-9]).*")
			match := reg.FindStringSubmatch(elem)
			if len(match) > 0 {
				num := MapStringToInt(match[1])
				day.output += num*10 + num
			}
		}
	}
	//part two

	day.output_snd = 0
	for _, elem := range day.input {
		///also matches the spelled out words
		reg := regexp.MustCompile("([1-9]|one|two|three|four|five|six|seven|eight|nine).*([1-9]|one|two|three|four|five|six|seven|eight|nine).*")
		match := reg.FindStringSubmatch(elem)
		if len(match) > 0 {

			num1 := getIntFromSpelledNumberOrVal(match[1])
			num2 := getIntFromSpelledNumberOrVal(match[2])
			day.output_snd += num1*10 + num2
		} else {
			reg := regexp.MustCompile("([1-9]|one|two|three|four|five|six|seven|eight|nine).*")
			if len(match) > 0 {
				match := reg.FindStringSubmatch(elem)
				num := MapStringToInt(match[1])
				day.output_snd += num*10 + num
			}
		}
	}
	fmt.Printf("day 1: pt1 : %d pt2: %d (currently incorrect)\n", day.output, day.output_snd)

}

func getIntFromSpelledNumberOrVal(inStr string) int {

	if len(inStr) == 1 {
		return MapStringToInt(inStr)
	} else {
		switch inStr {
		case "one":
			return 1
		case "two":
			return 2
		case "three":
			return 3
		case "four":
			return 4
		case "five":
			return 5
		case "six":
			return 6
		case "seven":
			return 7
		case "eight":
			return 8
		case "nine":
			return 9
		}
	}
	panic("should not have reached end of parse")
}
