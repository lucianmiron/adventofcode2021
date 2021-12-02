package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day2\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var horizontalTravel int = 0
	var depthTravel int = 0
	var aim int = 0
	var s []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = strings.Split(scanner.Text(), " ")
		var currentStep, _ = strconv.Atoi(s[1])
		switch s[0] {
		case "forward":
			horizontalTravel += currentStep
			depthTravel += aim * currentStep
		case "down":
			aim += currentStep
		case "up":
			aim -= currentStep
		}

	}
	fmt.Println("Horizontal x dept: ", horizontalTravel*depthTravel)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
