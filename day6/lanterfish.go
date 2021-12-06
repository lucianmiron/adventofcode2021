package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func notmain() {
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day6\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fishArray []int
	for scanner.Scan() {
		var fishes = strings.Split(scanner.Text(), ",")
		fishArray = make([]int, len(fishes))
		for i := 0; i < len(fishes); i++ {
			var fish, _ = strconv.Atoi(fishes[i])
			fishArray[i] = fish
		}
	}

	var countFishesForNextDay int = 0
	for day := 1; day <= 255; day++ {
		countFishesForNextDay = 0
		for i := 0; i < len(fishArray); i++ {
			fishArray[i]--
			if fishArray[i] == 0 {
				fishArray[i] = 7
				countFishesForNextDay++
			}
		}
		var fishesForNextDay = make([]int, countFishesForNextDay)
		for j := 0; j < len(fishesForNextDay); j++ {
			fishesForNextDay[j] = 9
		}
		fishArray = append(fishArray, fishesForNextDay...)
	}

	fmt.Printf("%d", len(fishArray))
}
