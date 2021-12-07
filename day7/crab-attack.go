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
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day7\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var horizontalPositions [5000]int
	var horizontalPositionsStr []string
	for scanner.Scan() {
		horizontalPositionsStr = strings.Split(scanner.Text(), ",")
	}

	for i := 0; i < len(horizontalPositionsStr); i++ {
		var horizontalPosition, _ = strconv.Atoi(horizontalPositionsStr[i])
		horizontalPositions[horizontalPosition]++
	}

	var lowestFuelHorizontalPos int = int(^uint(0) >> 1)
	var positionFuel [5000]int
	for i := 0; i < len(positionFuel); i++ {
		for j := 0; j < len(horizontalPositions); j++ {
			if horizontalPositions[j] != 0 {
				var individualFuel int = 0
				for k := 1; k <= int(math.Abs(float64(i)-float64(j))); k++ {
					individualFuel += k
				}
				positionFuel[i] += individualFuel * horizontalPositions[j]
			}
		}
		if lowestFuelHorizontalPos >= positionFuel[i] {
			lowestFuelHorizontalPos = positionFuel[i]
		}
	}
	fmt.Printf("Lowest horizontal position:%d", lowestFuelHorizontalPos)

}
