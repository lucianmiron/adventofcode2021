package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day3\\input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	var lineLength int
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if lineLength == 0 {
			lineLength = len(scanner.Text())
		}
	}
	var oxygenRatingBin = getOxygenRating(lines, 0)
	var co2RatingBin = getCo2Rating(lines, 0)
	var oxygenRatingBinInt, _ = strconv.Atoi(oxygenRatingBin)
	var co2RatingBinInt, _ = strconv.Atoi(co2RatingBin)

	var oxygenRating int
	var co2Rating int
	for i := 0; i < lineLength; i++ {
		oxygenRating += int(math.Pow(float64(2), float64(i))) * (oxygenRatingBinInt % 10)
		co2Rating += int(math.Pow(float64(2), float64(i))) * (co2RatingBinInt % 10)
		oxygenRatingBinInt = oxygenRatingBinInt / 10
		co2RatingBinInt = co2RatingBinInt / 10
	}

	fmt.Printf("Life Support:", oxygenRating*co2Rating)
}

func getOxygenRating(input []string, startIndex int) string {
	if len(input) == 1 {
		return input[0]
	}
	var zeroInput []string
	var oneInput []string
	for i := 0; i < len(input); i++ {
		if input[i][startIndex:startIndex+1] == "0" {
			zeroInput = append(zeroInput, input[i])
		} else {
			oneInput = append(oneInput, input[i])
		}
	}
	if len(oneInput) >= len(zeroInput) {
		return getOxygenRating(oneInput, startIndex+1)
	} else {
		return getOxygenRating(zeroInput, startIndex+1)
	}
}

func getCo2Rating(input []string, startIndex int) string {
	if len(input) == 1 {
		return input[0]
	}
	var zeroInput []string
	var oneInput []string
	for i := 0; i < len(input); i++ {
		if input[i][startIndex:startIndex+1] == "0" {
			zeroInput = append(zeroInput, input[i])
		} else {
			oneInput = append(oneInput, input[i])
		}
	}
	if len(oneInput) < len(zeroInput) && len(oneInput) > 0 {
		return getCo2Rating(oneInput, startIndex+1)
	} else {
		return getCo2Rating(zeroInput, startIndex+1)
	}
}
