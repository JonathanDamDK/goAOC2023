package main

import (
	"bufio"
	"log"
	"os"
)


func MapFileToStringArr(inputPath string) []string {
	//Open file
	var lines []string
	file, err := os.Open(inputPath)
	//The file is not found
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines

}
