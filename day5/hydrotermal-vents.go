package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day5\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var vents [1000][1000]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var coordinates = strings.Split(scanner.Text(), " -> ")
		var leftCoordinates = strings.Split(coordinates[0], ",")
		var x1, _ = strconv.Atoi(leftCoordinates[0])
		var y1, _ = strconv.Atoi(leftCoordinates[1])
		var rightCoordinates = strings.Split(coordinates[1], ",")
		var x2, _ = strconv.Atoi(rightCoordinates[0])
		var y2, _ = strconv.Atoi(rightCoordinates[1])

		var startFrom int
		var endTo int
		var startFromY int
		var endToY int
		var diagonalIncrement int
		if x1 == x2 {
			if y1 < y2 {
				startFrom = y1
				endTo = y2
			} else {
				startFrom = y2
				endTo = y1
			}
			for i := startFrom; i <= endTo; i++ {
				vents[i][x1]++
			}
		} else if y1 == y2 {
			if x1 < x2 {
				startFrom = x1
				endTo = x2
			} else {
				startFrom = x2
				endTo = x1
			}
			for i := startFrom; i <= endTo; i++ {
				vents[y1][i]++
			}
		} else if math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)) { //diagonal always
			if y1 < y2 {
				startFromY = y1
				startFrom = x1
				endToY = y2
				endTo = x2
			} else {
				startFromY = y2
				startFrom = x2
				endToY = y1
				endTo = x1
			}
			if startFrom < endTo {
				diagonalIncrement = 1
			} else {
				diagonalIncrement = -1
			}
			for startFrom != endTo+diagonalIncrement && startFromY != endToY+1 {
				vents[startFromY][startFrom]++
				startFromY++
				startFrom = startFrom + diagonalIncrement
			}
		}
	}

	var countVents int = 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if vents[i][j] >= 2 {
				countVents++
			}
		}
	}
	fmt.Printf("Vents number:%d", countVents)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
