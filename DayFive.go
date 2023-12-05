package main

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	seed_soil            int = 0
	soil_fertilizer          = 1
	fertilizer_water         = 2
	water_light              = 3
	light_temperature        = 4
	temperatire_humidity     = 5
	humidity_location        = 6
)

type DayFive struct {
	inputPath string
	input     []string
	outputpt1 int
	outputpt2 int
}

type inOutrange struct {
	inputStart  int
	outputStart int
	numRange    int
}

//if(input < inputStart + numRange) output = outputStart + input - inputStart

func (day DayFive) Solve() {
	part1(&day)
	part2(&day)
	fmt.Printf("day 5 part 1: %d day 5 part 2: %d\n", day.outputpt1, day.outputpt2)
}
func part1(day *DayFive) {
	lines := MapFileToStringArr(day.inputPath)
	var rangeArr [humidity_location + 1][]inOutrange
	var seeds []int
	globalRangeIndex := -1
	for _, line := range lines {
		count := 0
		regex := regexp.MustCompile(" ?(\\d+)")
		matches := regex.FindAllStringSubmatch(line, -1)
		var rangeElement inOutrange
		for _, match := range matches {
			num, _ := strconv.Atoi(match[1])
			//if the current index is -1 it means that we are parsing seeds from the input thus we should not alter the rangeElement
			if globalRangeIndex > -1 {
				switch count {
				case 0:
					rangeElement.outputStart = num
					break
				case 1:
					rangeElement.inputStart = num
					break
				case 2:
					rangeElement.numRange = num
					break
				}
				count++
			} else {
				seeds = append(seeds, num)
			}
		}
		//new range starts in input line
		if len(matches) == 0 && len(line) > 0 {
			globalRangeIndex += 1
		} else {
			//only if there is charactes in a line, append rangeElement
			if globalRangeIndex >= 0 && len(line) > 0 {
				rangeArr[globalRangeIndex] = append(rangeArr[globalRangeIndex], rangeElement)
			}
		}
	}
	locations := make([]int, len(seeds))
	//from here input has been parsed so we loop through each number to get its position
	for i, seed := range seeds {
		currValue := seed
		for _, rangeCollection := range rangeArr {
			for _, rangeObj := range rangeCollection {
				if currValue >= rangeObj.inputStart && currValue <= (rangeObj.inputStart-1+rangeObj.numRange) {
					currValue = rangeObj.outputStart + currValue - rangeObj.inputStart
					break
				}
			}

		}
		locations[i] = currValue
	}
	day.outputpt1 = min(locations)

}

func part2(day *DayFive) {
	day.outputpt2 = 1000000000000000000
	lines := MapFileToStringArr(day.inputPath)
	var rangeArr [humidity_location + 1][]inOutrange
	var seeds []int
	globalRangeIndex := -1
	for _, line := range lines {
		count := 0
		regex := regexp.MustCompile(" ?(\\d+)")
		matches := regex.FindAllStringSubmatch(line, -1)
		var rangeElement inOutrange
		for _, match := range matches {
			num, _ := strconv.Atoi(match[1])
			//if the current index is -1 it means that we are parsing seeds from the input thus we should not alter the rangeElement
			if globalRangeIndex > -1 {
				switch count {
				case 0:
					rangeElement.outputStart = num
					break
				case 1:
					rangeElement.inputStart = num
					break
				case 2:
					rangeElement.numRange = num
					break
				}
				count++
			} else {
				seeds = append(seeds, num)
			}
		}
		//new range starts in input line
		if len(matches) == 0 && len(line) > 0 {
			globalRangeIndex += 1
		} else {
			//only if there is charactes in a line, append rangeElement
			if globalRangeIndex >= 0 && len(line) > 0 {
				rangeArr[globalRangeIndex] = append(rangeArr[globalRangeIndex], rangeElement)
			}
		}
	}
	minvalue := 100000000000000000
	seedPairs := make([]inOutrange, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		seedIndex := 0
		if i != 0 {
			seedIndex = i / 2
		}
		//fmt.Printf("seed: %d %d \n", seeds[i], seeds[i+1])
		seedPairs[seedIndex] = inOutrange{seeds[i], 0, seeds[i+1]}
	}
	var resultArr [][]inOutrange
	for _, seedPair := range seedPairs {
		currEvaluationRanges := []inOutrange{seedPair}
		for _, rangeCollection := range rangeArr {
			for evalIdx, evalRange := range currEvaluationRanges {
				findPairingRange(&currEvaluationRanges, rangeCollection, evalRange, evalIdx)
			}
		}

		resultArr = append(resultArr, currEvaluationRanges)
		
		minvalue = min_two(minvalue,min_row(currEvaluationRanges))

	}
	day.outputpt2 = minvalue

}
func min_two(num1 int, num2 int) int{
	if num1 < num2 {
			return num1
	} else {
		return num2
	}
}
func min_row(evalRange []inOutrange) int {
	minElem := 1000000000000000000
	for _, row := range evalRange {
		if row.inputStart < minElem {
			minElem = row.inputStart
		}
	}
	return minElem
}

func min(arr []int) int {
	minElem := 1000000000000000000
	for _, num := range arr {
		if num < minElem {
			minElem = num
		}
	}
	return minElem
}

func findPairingRange(currElements *[]inOutrange, rangeCollection []inOutrange, evalRange inOutrange, evalIdx int) {
	foundWholeRange := false
	for _, rangeObj := range rangeCollection {
		if evalRange.inputStart >= rangeObj.inputStart && evalRange.inputStart+evalRange.numRange <= rangeObj.inputStart+rangeObj.numRange {
			(*currElements)[evalIdx].inputStart = rangeObj.outputStart + evalRange.inputStart - rangeObj.inputStart
			foundWholeRange = true
		}
	}
	if !foundWholeRange {
		for _, rangeObj := range rangeCollection {
			//calculate values for easy readability	
			evalStart := evalRange.inputStart
			evalEnd := evalRange.inputStart + evalRange.numRange
			rangeStart := rangeObj.inputStart
			rangeEnd := rangeObj.numRange + rangeObj.inputStart
			//if we have a range where there are a number of beginning elements in evalRange that does not fit, we split
			if evalStart < rangeStart && evalEnd > rangeStart && evalEnd <= rangeEnd {
				splitObj := inOutrange{}
				originalRange := evalRange.numRange
				//the curr range is now set to the elements between evalStart and rangeStart, that is the elements that do not fit into the range object
				(*currElements)[evalIdx].numRange = rangeStart - evalStart
				(*currElements)[evalIdx].inputStart = evalStart 
				//splitObj represents all the elements that do fit, thus we set the numRange to the remaining numbers
				splitObj.numRange = originalRange - (*currElements)[evalIdx].numRange
				//we know that splitObj without map would start at rangeob.inputStart since we want to map we set it imidiatly to the outputStart
				splitObj.inputStart = rangeObj.outputStart 	
				(*currElements) = append((*currElements), splitObj)
				//our (*currElements)[evalIdx] needs to be rechecked to see if it fits in to any range now that it has been altered
				findPairingRange(currElements, rangeCollection, (*currElements)[evalIdx], evalIdx)

				//there are more elements in our evalRange than in the range we are checking, thus a split is required
			} else if evalStart >= rangeStart && evalStart < rangeEnd && evalEnd > rangeEnd {
				splitObj := inOutrange{}
				//currElements[evalIdx] is set to represent the values that are present in evalRange thats why we take all elements up to the rangeEnd
				(*currElements)[evalIdx].numRange = rangeEnd - evalStart
				//we then take all elements 
				splitObj.numRange = evalEnd - rangeEnd
				splitObj.inputStart = rangeEnd
				//we set the start equal to the output + the offset in the range ie if eval is starting at 77 and range at 73, and output start is 5 we set inputStart = 5 + 77-73 = 9
				(*currElements)[evalIdx].inputStart = rangeObj.outputStart + evalStart - 	rangeStart
				*currElements = append((*currElements), splitObj)
				//since splitObj now represents a new range the new range is also checked to see if it has any matches
				findPairingRange(currElements, rangeCollection, splitObj, len(*currElements)-1)
			}
		}
	}

}
