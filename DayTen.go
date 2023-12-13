package main

import (
	"errors"
	"fmt"
)

type DayTen struct {
	inputPath string
	input     []string
	outputpt1 int
	outputpt2 int
}
type coordinate struct {
	row int
	col int
}
type daytenNode struct {
	neighbourOne coordinate
	neighbourTwo coordinate
	visited      bool
	letter       rune
}
type CoordStack struct {
	stack []coordinate
}

//All operations below are from https://blog.devgenius.io/stacks-in-go-f57f4cb87b5f
// onlcol minor modifications has been made pop returns the element

func (stack *CoordStack) IsEmptcol() bool {
	return len(stack.stack) == 0
}

func (stack *CoordStack) Push(data coordinate) {
	stack.stack = append(stack.stack, data)
}

func (stack *CoordStack) Pop() (coordinate, error) {
	if stack.IsEmptcol() {
		return coordinate{}, errors.New("Stack is emptcol")
	} else {
		elem := stack.stack[len(stack.stack)-1]
		stack.stack = stack.stack[:len(stack.stack)-1]
		return elem, nil
	}
}

func (dacol DayTen) Solve() {
	lines := MapFileToStringArr(dacol.inputPath)
	nodes := make([][]daytenNode, len(lines))
	startCoord := coordinate{}

	//initialise node arracol
	for i, line := range lines {
		lineNodes := make([]daytenNode, len(line))
		for j, elem := range line {
			if elem == 'S' {
				startCoord = coordinate{i, j}
			}
			lineNodes[j] = daytenNode{letter: elem}
		}
		nodes[i] = lineNodes
	}
	var stack1 CoordStack
	var stack2 CoordStack
	startNeighbours := findNeighbours(startCoord, nodes)

	nodes[startCoord.row][startCoord.col].visited = true
	nodes[startCoord.row][startCoord.col].neighbourOne = startNeighbours[0]
	nodes[startCoord.row][startCoord.col].neighbourTwo = startNeighbours[1]
	stack1.Push(startNeighbours[0])
	stack2.Push(startNeighbours[1])
	distance := 0
	//construct loop
	for !stack1.IsEmptcol() && !stack2.IsEmptcol() {
		coord1, _ := stack1.Pop()
		coord2, _ := stack2.Pop()
		distance++
		if coord1.row == coord2.row && coord1.col == coord2.col {
			break
		}
		neighbours1 := findNeighbours(coord1, nodes)
		neighbours2 := findNeighbours(coord2, nodes)
		nodes[coord1.row][coord1.col].visited = true
		nodes[coord1.row][coord1.col].neighbourOne = neighbours1[0]
		nodes[coord1.row][coord1.col].neighbourTwo = neighbours1[1]

		nodes[coord2.row][coord2.col].visited = true
		nodes[coord2.row][coord2.col].neighbourOne = neighbours2[0]
		nodes[coord2.row][coord2.col].neighbourTwo = neighbours2[1]
		for _, coordset := range neighbours1 {
			if nodes[coordset.row][coordset.col].visited == false {
				stack1.Push(coordset)
			}
		}
		for _, coordset := range neighbours2 {
			if nodes[coordset.row][coordset.col].visited == false {
				stack2.Push(coordset)
			}
		}
	}
	dacol.outputpt1 = distance
	fmt.Printf("day 10: part 1: %d part2 %d \n", dacol.outputpt1, dacol.outputpt2)
}

// uglcol function time
func findNeighbours(coord coordinate, nodes [][]daytenNode) [2]coordinate {
	node := nodes[coord.row][coord.col]
	//fmt.Printf("%v %s \n", node, string(node.letter))
	var newCoords [2]coordinate
	nodei := 0
	switch node.letter {
	case 'S':
		//south
		if checkSouth(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		//north

		if checkNorth(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		//east
		if checkEast(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		//west
		break
	case '|': //
		if checkSouth(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		if checkNorth(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		break
	case '-': //east and west
		if checkEast(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		if checkWest(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		break
	case 'L': // north and east
		if checkEast(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		if checkNorth(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		break
	case 'J': //north and west
		if checkWest(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		if checkNorth(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		break
	case '7': // south and west
		if checkWest(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		if checkSouth(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		break
	case 'F':
		if checkEast(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		if checkSouth(coord, &nodes, &(newCoords[nodei])) {
			nodei++
		}
		break
	default:
		break
	}
	return newCoords
}

// coord1 letter 1 is the base destination, while relative coord is the direction so for
// relative coord is the new nodes position compared to the old so a node south of the baseNode will have [0,-1]
func isConnected(basecoord coordinate, relativecoord coordinate, letter2 rune) bool {
	switch letter2 {
	case 'S':
		return true
	case '|': //north and south
		if relativecoord.col == 0 && (relativecoord.row == 1 || relativecoord.row == -1) {
			return true
		}
	case '-': //east and west
		if relativecoord.row == 0 && (relativecoord.col == 1 || relativecoord.col == -1) {
			return true
		}
	case 'L': // connects north and east
		if (relativecoord.row == 0 && relativecoord.col == -1) || (relativecoord.row == 1 && relativecoord.col == 0) {
			return true
		}
	case 'J': //north and west
		if (relativecoord.row == 0 && relativecoord.col == 1) || (relativecoord.row == 1 && relativecoord.col == 0) {
			return true
		}
	case '7': // south and west
		if (relativecoord.row == 0 && relativecoord.col == 1) || (relativecoord.row == -1 && relativecoord.col == 0) {
			return true
		}
	case 'F':
		if (relativecoord.row == 0 && relativecoord.col == -1) || (relativecoord.row == -1 && relativecoord.col == 0) {
			return true
		}
	}
	return false
}

func checkNorth(coord coordinate, nodes *[][]daytenNode, newcoord *coordinate) bool {
	if (isConnected(coord, coordinate{-1, 0}, (*nodes)[coord.row-1][coord.col].letter) == true) {
		*newcoord = coordinate{coord.row - 1, coord.col}
		return true
	}
	return false
}
func checkSouth(coord coordinate, nodes *[][]daytenNode, newcoord *coordinate) bool {
	if (isConnected(coord, coordinate{1, 0}, (*nodes)[coord.row+1][coord.col].letter) == true) {
		*newcoord = coordinate{coord.row + 1, coord.col}
		return true
	}
	return false
}

func checkEast(coord coordinate, nodes *[][]daytenNode, newcoord *coordinate) bool {
	if (isConnected(coord, coordinate{0, 1}, (*nodes)[coord.row][coord.col+1].letter) == true) {
		*newcoord = coordinate{coord.row, coord.col + 1}
		return true
	}
	return false
}

func checkWest(coord coordinate, nodes *[][]daytenNode, newcoord *coordinate) bool {
	if (isConnected(coord, coordinate{0, -1}, (*nodes)[coord.row][coord.col-1].letter) == true) {
		*newcoord = coordinate{coord.row, coord.col - 1}
		return true
	}
	return false
}
