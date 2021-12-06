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
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day6\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fishArray [9]int
	for scanner.Scan() {
		var fishes = strings.Split(scanner.Text(), ",")
		for i := 0; i < len(fishes); i++ {
			var fish, _ = strconv.Atoi(fishes[i])
			fishArray[fish]++
		}
	}

	var fishesForCurrentDay [9]int
	var fishesBornYesterday int = 0
	for day := 1; day <= 256; day++ {
		for i := 0; i <= 6; i++ {
			var index int = 0
			index = (i - 1) % 7
			if index < 0 {
				index += 7
			}
			fishesForCurrentDay[index] = fishArray[i]
		}
		fishesForCurrentDay[8] = fishesBornYesterday
		fishesBornYesterday = fishesForCurrentDay[0]
		fishesForCurrentDay[6] += fishArray[7]
		fishesForCurrentDay[7] = fishArray[8]

		fishArray = fishesForCurrentDay
	}

	var totalNoOfFish int = 0
	for i := 0; i <= 8; i++ {
		totalNoOfFish += fishArray[i]
	}
	fmt.Printf("%d", totalNoOfFish)
}
