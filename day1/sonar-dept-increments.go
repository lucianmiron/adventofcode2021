package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	const MaxInt = int(^uint(0) >> 1)
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day1\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var previousMeasurement int = int(MaxInt)
	var depthIncrements int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var currentMeasurement, _ = strconv.Atoi(scanner.Text())
		if currentMeasurement > previousMeasurement {
			depthIncrements++
		}
		previousMeasurement = currentMeasurement
	}
	fmt.Println("Total number of depth increments are: ", depthIncrements)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
