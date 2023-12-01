package main

import (
	"strconv"
)

func MapStringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return num
}
