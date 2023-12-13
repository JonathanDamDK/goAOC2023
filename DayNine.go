package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type DayNine struct {
	inputPath string
	input     []string
	outputpt1 int
	outputpt2 int
}

type Stack struct {
	stack [][]int
}

//All operations below are from https://blog.devgenius.io/stacks-in-go-f57f4cb87b5f
// only minor modifications has been made pop returns the element

func (stack *Stack) Push(data []int) {
	stack.stack = append(stack.stack, data)
}

func (stack *Stack) Pop() ([]int, error) {
	if stack.IsEmpty() {
		return make([]int, 0), errors.New("Stack is empty")
	} else {
		elem := stack.stack[len(stack.stack)-1]
		stack.stack = stack.stack[:len(stack.stack)-1]
		return elem, nil
	}
}

func (stack *Stack) Peek() []int {
	return stack.stack[len(stack.stack)-1]
}
func (stack *Stack) IsEmpty() bool {
	if len(stack.stack) == 0 {
		return true
	}
	return false
}

func (day DayNine) Solve() {
	lines := MapFileToStringArr(day.inputPath)
	var readings [][]int
	//read input
	for _, line := range lines {
		expr := regexp.MustCompile(" *(-?\\d+)")
		match := expr.FindAllStringSubmatch(line, -1)
		arr := make([]int, len(match))
		for j, numMatch := range match {
			num, _ := strconv.Atoi(numMatch[1])
			arr[j] = num
		}
		readings = append(readings, arr)
	}

	totalPrediction := 0
	totalPredictionpart2 := 0
	for _, reading := range readings {
		var stack Stack
		stack.Push(reading)
		curr := stack.Peek()
		//Create stack
		for  !allZero(curr){
			arr := make([]int, len(curr)-1)
			for i := 0; i < len(curr)-1; i++ {
				num := curr[i+1] - curr[i]
				arr[i] = num
			}
			stack.Push(arr)
			curr = arr
		}
		currDifference := 0
		currDifferencepart2 := 0
		for !stack.IsEmpty() {
			elem, err := stack.Pop()
			if err != nil {
				println(err)
			}
			currDifference = elem[len(elem)-1] + currDifference
			currDifferencepart2 = elem[0] - currDifferencepart2
		}

		totalPrediction += currDifference
		totalPredictionpart2 += currDifferencepart2
		//fmt.Printf("%v %d \n", stack, currDifference)
	}
	day.outputpt1 = totalPrediction
	day.outputpt2 = totalPredictionpart2
	fmt.Printf("day 9: part 1: %d part 2: %d \n", day.outputpt1, day.outputpt2)
}

func allZero(arr []int) bool{
	res := true 
	for _, elem := range arr {
		if(elem != 0){
			res = false 
			break 
		}
	}
	return res
}
