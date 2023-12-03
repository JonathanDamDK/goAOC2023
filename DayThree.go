package main

import (
	"fmt"
	"strconv"
)

type DayThree struct {
	inputPath string
	input     []string
	outputpt1 int
	outputpt2 int
}
type digit struct {
	value       int
	rowIndex    int
	columnStart int
	columnEnd   int
	isPart      bool
}

type answerMatrixElem struct {
	typeNum  int
	intValue int
	symbol   rune
	visited  bool
}

func (day DayThree) Solve() {
	lines := MapFileToStringArr(day.inputPath)
	var answerMatrix [][]answerMatrixElem
	for rowIndex, line := range lines {
		parseLine(line, rowIndex, &answerMatrix)
	}
	partOne(&day, answerMatrix)
	partTwo(&day, answerMatrix)
	fmt.Printf("Day 3: part 1: %d part two : %d\n", day.outputpt1, day.outputpt2)

}
func partOne(day *DayThree, answerMatrix [][]answerMatrixElem) {
	for rowIndex, row := range answerMatrix {
		for colIndex, elem := range row {
			//if the current element is a number
			if elem.typeNum == 2 {
				//if it hasnt already been added to the total
				if elem.visited != true {
					var startCol, endCol, startRow, endRow int
					//get a 3x3 area acounting for edges to avoid overflow
					getIndices(&startRow, &endRow, &startCol, &endCol, colIndex, rowIndex, len(row), len(answerMatrix))
					//iterate over neighbours
					for i := startRow; i <= endRow; i++ {
						for j := startCol; j <= endCol; j++ {
							//if there is a symbol nearby
							if answerMatrix[i][j].typeNum == 1 {
								currIndex := colIndex
								//find the right most number, in order to get base ten representation
								for currIndex+1 < len(row) && answerMatrix[rowIndex][currIndex+1].typeNum == 2 {
									currIndex++
								} //right most number found
								//fmt.Printf("rightmostvalue : %d row : %d col : %d \n", answerMatrix[rowIndex][currIndex].intValue, rowIndex , currIndex)
								pow := 0
								val := 0
								//now go right to left marking all numbers as visited
								for currIndex >= 0 && answerMatrix[rowIndex][currIndex].typeNum == 2 {
									powval := 1
									//go has no power operation for integers thus we create one here
									for i := 1; i <= pow; i++ {
										powval *= 10
									}
									//add the number to the base ten representation
									val += powval * answerMatrix[rowIndex][currIndex].intValue
									pow += 1
									//mark it as visited and step left
									answerMatrix[rowIndex][currIndex].visited = true
									currIndex--
								}
								//the numeric value is now added to the output
								day.outputpt1 += val
								break
							}
						}
					}
				}
			}
		}
	}
}

func partTwo(day *DayThree, answerMatrix [][]answerMatrixElem) {
	for rowIndex, row := range answerMatrix {
		for colIndex, elem := range row {
			if elem.symbol == '*' {
				var startCol, endCol, startRow, endRow int
				//get a 3x3 area acounting for edges to avoid overflow
				getIndices(&startRow, &endRow, &startCol, &endCol, colIndex, rowIndex, len(row), len(answerMatrix))
				//iterate over neighbours
				numberNeighbours := 0
				neighbourTotal := 1
				for i := startRow; i <= endRow; i++ {
					parsingNumber := false // value to represent if the previous element was a number
					for j := startCol; j <= endCol; j++ {
						if answerMatrix[i][j].typeNum == 2 { // number
							if !parsingNumber {
								numIndexCol := j
								numStartRow := i
								//parse number start by finding right most element
								for numIndexCol < len(row)-1 && answerMatrix[numStartRow][numIndexCol+1].typeNum == 2 {
									numIndexCol += 1
								} // found right most element
								pow := 0
								val := 0
								//now go right to left in order to retrieve numerical value
								for numIndexCol >= 0 && answerMatrix[numStartRow][numIndexCol].typeNum == 2 {
									powval := 1
									//go has no power operation for integers thus we create one here
									for i := 1; i <= pow; i++ {
										powval *= 10
									}
									//add the number to the base ten representation
									val += powval * answerMatrix[numStartRow][numIndexCol].intValue
									pow += 1
									numIndexCol--
								}
								numberNeighbours += 1
								parsingNumber = true
								neighbourTotal *= val
							}
						} else { //
							parsingNumber = false
						}
					}
				}
				if numberNeighbours == 2 {
					day.outputpt2 += neighbourTotal
				}
			}
		}
	}
}
func parseLine(line string, rowIndex int, answerMatrix *[][]answerMatrixElem) int {
	stringIndex := 0
	var row = make([]answerMatrixElem, len(line))
	for stringIndex < len(line) {
		switch line[stringIndex] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num, _ := strconv.Atoi(string(line[stringIndex]))
			row[stringIndex] = answerMatrixElem{2, num, 'n', false}
			stringIndex++
			break
		case '.':
			stringIndex++
			break
		default: // symbol
			row[stringIndex] = answerMatrixElem{1, 0, rune(line[stringIndex]), false}
			stringIndex++
			break
		}
	}
	*answerMatrix = append(*answerMatrix, row)
	return 1
}

// /sets the appropriate values for a 3x3 search vasinity
func getIndices(startRow *int, endRow *int, startCol *int, endCol *int, colIndex int, rowIndex int, rowLen int, matrixLen int) {
	if rowIndex > 0 {
		*startRow = rowIndex - 1 //look at rows above
	} else {
		*startRow = 0 // first row
	}

	if rowIndex < matrixLen-1 {
		*endRow = rowIndex + 1
	} else {
		*endRow = rowIndex //evaluating the last row
	}

	if colIndex > 0 {
		*startCol = colIndex - 1
	} else {
		*startCol = 0
	}
	if colIndex < rowLen-1 {
		*endCol = colIndex + 1
	} else {
		*endCol = colIndex
	}
}
