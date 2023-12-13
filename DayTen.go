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
	x int
	y int
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
// only minor modifications has been made pop returns the element

func (stack *CoordStack) IsEmpty() bool {
	return len(stack.stack) == 0
}

func (stack *CoordStack) Push(data coordinate) {
	stack.stack = append(stack.stack, data)
}

func (stack *CoordStack) Pop() (coordinate, error) {
	if stack.IsEmpty() {
		return coordinate{}, errors.New("Stack is empty")
	} else {
		elem := stack.stack[len(stack.stack)-1]
		stack.stack = stack.stack[:len(stack.stack)-1]
		return elem, nil
	}
}

func (day DayTen) Solve() {
	lines := MapFileToStringArr(day.inputPath)
	nodes := make([][]daytenNode, len(lines))
	startCoord := coordinate{}

	//initialise node array
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

	nodes[startCoord.x][startCoord.y].visited = true
	nodes[startCoord.x][startCoord.y].neighbourOne = startNeighbours[0]
	nodes[startCoord.x][startCoord.y].neighbourTwo = startNeighbours[1]
	stack1.Push(startNeighbours[0])
	stack2.Push(startNeighbours[1])
	fmt.Printf("%v\n", startCoord)
	distance := 0
	//construct loop
	for !stack1.IsEmpty() && !stack2.IsEmpty() {
		coord1, _ := stack1.Pop()
		coord2, _ := stack2.Pop()
		distance++
		fmt.Printf("coord1: %v %s, coord2 : %v%s\n", coord1, string(nodes[coord1.x][coord1.y].letter), coord2, string(nodes[coord2.x][coord2.y].letter))
		if coord1.x == coord2.x && coord1.y == coord2.y {
			break
		}
		neighbours1 := findNeighbours(coord1, nodes)
		neighbours2 := findNeighbours(coord2, nodes)
		nodes[coord1.x][coord1.y].visited = true
		nodes[coord1.x][coord1.y].neighbourOne = neighbours1[0]
		nodes[coord1.x][coord1.y].neighbourTwo = neighbours1[1]

		nodes[coord2.x][coord2.y].visited = true
		nodes[coord2.y][coord2.y].neighbourOne = neighbours2[0]
		nodes[coord2.y][coord2.y].neighbourTwo = neighbours2[1]
		for _, coordset := range neighbours1 {
			if coordset.x == 0 && coordset.y == 0 {
				fmt.Printf("Zero added: %v  %v \n", string(nodes[coord2.x][coord2.y].letter), neighbours1)
			}
			if nodes[coordset.x][coordset.y].visited == false {
				stack1.Push(coordset)
			}
		}
		for _, coordset := range neighbours2 {
			if nodes[coordset.x][coordset.y].visited == false {
				stack2.Push(coordset)
			}
		}
	}
	day.outputpt1 = distance
	fmt.Printf("day 10: part 1: %d part2 %d \n", day.outputpt1, day.outputpt2)
}

// ugly function time
func findNeighbours(coord coordinate, nodes [][]daytenNode) [2]coordinate {
	node := nodes[coord.y][coord.x]
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
	case '|': //
		if relativecoord.x == 0 && (relativecoord.y == 1 || relativecoord.y == -1) {
			return true
		}
	case '-': //east and west
		if relativecoord.y == 0 && (relativecoord.x == 1 || relativecoord.x == -1) {
			return true
		}
	case 'L': // north and east
		if (relativecoord.x == 0 && relativecoord.y == -1) || (relativecoord.x == -1 && relativecoord.y == 0) {
			return true
		}
	case 'J': //north and west
		if (relativecoord.x == 0 && relativecoord.y == -1) || (relativecoord.x == 1 && relativecoord.y == 0) {
			return true
		}
	case '7': // south and west
		if (relativecoord.x == 0 && relativecoord.y == 1) || (relativecoord.x == 1 && relativecoord.y == 0) {
			return true
		}
	case 'F':
		if (relativecoord.x == 0 && relativecoord.y == -1) || (relativecoord.x == -1 && relativecoord.y == 0) {
			return true
		}
	}
	return false
}

func checkNorth(coord coordinate, nodes *[][]daytenNode, newcoord *coordinate) bool {
	if (isConnected(coord, coordinate{0, 1}, (*nodes)[coord.x][coord.y-1].letter) == true) {
		*newcoord = coordinate{coord.x, coord.y - 1}
		return true
	}
	return false
}
func checkSouth(coord coordinate, nodes *[][]daytenNode, newcoord *coordinate) bool {
	if (isConnected(coord, coordinate{0, -1}, (*nodes)[coord.y+1][coord.x].letter) == true) {
		*newcoord = coordinate{coord.x, coord.y + 1}
		return true
	}
	return false
}

func checkEast(coord coordinate, nodes *[][]daytenNode, newcoord *coordinate) bool {
	if (isConnected(coord, coordinate{1, 0}, (*nodes)[coord.y][coord.x+1].letter) == true) {
		*newcoord = coordinate{coord.x + 1, coord.y}
		return true
	}
	return false
}

func checkWest(coord coordinate, nodes *[][]daytenNode, newcoord *coordinate) bool {

	if (isConnected(coord, coordinate{-1, 0}, (*nodes)[coord.y][coord.x-1].letter) == true) {
		*newcoord = coordinate{coord.x - 1, coord.y}
		return true
	}
	return false
}
