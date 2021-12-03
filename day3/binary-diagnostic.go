package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func binarydiagnostic() {
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day3\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Count how many 1's there are on each position
	var countOnes []int
	var fileEntries int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Bytes()
		for i := 0; i < len(line); i++ {
			if len(countOnes) <= i {
				countOnes = append(countOnes, 0)
			}
			if line[i] == 49 { // 1
				countOnes[i]++
			}
		}
		fileEntries++
	}

	// Build binary versions of gama and epsilon
	var gamaStr string = ""
	var epsilonStr string = ""
	for i := 0; i < len(countOnes); i++ {
		if countOnes[i] > fileEntries/2 {
			gamaStr += "1"
			epsilonStr += "0"
		} else {
			gamaStr += "0"
			epsilonStr += "1"
		}
	}

	// Convert from binary to decimal
	var gamaInt, _ = strconv.Atoi(gamaStr)
	var epsilonInt, _ = strconv.Atoi(epsilonStr)
	var gama int = 0
	var epsilon int = 0
	for i := 0; i < len(countOnes); i++ {
		gama += int(math.Pow(float64(2), float64(i))) * (gamaInt % 10)
		epsilon += int(math.Pow(float64(2), float64(i))) * (epsilonInt % 10)
		gamaInt = gamaInt / 10
		epsilonInt = epsilonInt / 10
	}
	var comsumntion int = gama * epsilon
	fmt.Println("Submarine consumption:", comsumntion)
}
