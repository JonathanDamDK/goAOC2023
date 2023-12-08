package main

import (
	"fmt"
	"regexp"
)

type DayEight struct {
	inputPath string
	input     []string
	outputpt1 int
	outputpt2 int
}

type mapElement struct {
	left  string
	right string
}

func (day DayEight) Solve() {
	lines := MapFileToStringArr(day.inputPath)
	directions := ""
	mapLR := make(map[string]mapElement)
	for i, line := range lines {
		if i == 0 {
			reg := regexp.MustCompile("([LR]*)")
			match := reg.FindStringSubmatch(line)
			directions = match[1]
		} else {
			reg := regexp.MustCompile(`([1-9]*[A-Z]*) *= *\(([1-9]*[A-Z]*) *, *([1-9]*[A-Z]*)\)`)
			match := reg.FindStringSubmatch(line)
			if len(match) > 0 {
				mapLR[match[1]] = mapElement{match[2], match[3]}
			}
		}

		//println(line)
	}
	found := false
	curr := "AAA"
	count := 0

	//part 1
	for !found {
		for i := 0; i < len(directions); i++ {
			if directions[i] == 'L' {
				curr = mapLR[curr].left

			} else {
				curr = mapLR[curr].right
			}
			count++
			if curr == "ZZZ" {
				found = true
			}
		}
	}
	day.outputpt1 = count

	//part 2 gave up and used reddit to discover that i can use LCM rest is my own design
	var startNodes []string
	for key := range mapLR {
		//get all nodes that end with A
		if key[len(key)-1] == 'A' {
			startNodes = append(startNodes, key)
		}
	}
	pathArr := make([]int, len(startNodes))

	//find every nodes shortest path to a node that ends with Z
	for i, node := range startNodes {
		found = false
		count = 0
		curr = node
		for !found {
			for i := 0; i < len(directions); i++ {
				if directions[i] == 'L' {
					curr = mapLR[curr].left
					count++
				} else {
					curr = mapLR[curr].right
					count++
				}
				if curr[len(curr)-1] == 'Z' {
					found = true
				} else {
					found = false
				}
			}
		}
		pathArr[i] = count
	}
	count = 0
	//use LCM 
	day.outputpt2 = lcm(pathArr, len(pathArr))
	fmt.Printf("Day 8: part 1 : %d part 2 : %d\n", day.outputpt1, day.outputpt2)
}

// source https://www.geeksforgeeks.org/lcm-of-given-array-elements/
func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// lowest common multiplier
func lcm(arr []int, n int) int {
	ans := arr[0]

	for i := 1; i < n; i++ {
		ans = (ans * arr[i]) / (gcd(arr[i], ans))
	}
	return ans
}
