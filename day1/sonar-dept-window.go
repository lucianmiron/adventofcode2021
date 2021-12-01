package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day1\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var depthIncrements int = 0
	var index int = 0
	var firstW int = 0
	var secondW int = 0
	var thirdW int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var currentMeasurement, _ = strconv.Atoi(scanner.Text())

		switch index % 3 {
		case 0:
			secondW = secondW + currentMeasurement
			thirdW = thirdW + currentMeasurement
			if index > 3 && firstW < secondW {
				depthIncrements++
			}
			firstW = currentMeasurement
		case 1:
			firstW = firstW + currentMeasurement
			thirdW = thirdW + currentMeasurement
			if index > 3 && secondW < thirdW {
				depthIncrements++
			}
			secondW = currentMeasurement
		case 2:
			firstW = firstW + currentMeasurement
			secondW = secondW + currentMeasurement
			if index > 3 && thirdW < firstW {
				depthIncrements++
			}
			thirdW = currentMeasurement
		}

		index++
	}
	fmt.Println("Total number of depth increments are: ", depthIncrements)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
