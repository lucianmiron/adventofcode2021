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
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day13\\input-folds.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var x, y int
	for scanner.Scan() {
		fold := strings.Split(scanner.Text(), "=")
		if fold[0] == "y" {
			y, _ = strconv.Atoi(fold[1])
		} else {
			x, _ = strconv.Atoi(fold[1])
		}
		if y != 0 && x != 0 {
			break
		}
	}
	y = y*2 + 1
	x = x*2 + 1
	points := make([][]int, y)
	for i := range points {
		points[i] = make([]int, x)
	}

	file2, err2 := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day13\\input-points.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer file2.Close()

	scanner = bufio.NewScanner(file2)
	for scanner.Scan() {
		coordinates := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])

		points[y][x] = 1
	}

	file3, err3 := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day13\\input-folds.txt")
	if err3 != nil {
		log.Fatal(err3)
	}

	scanner2 := bufio.NewScanner(file3)
	for scanner2.Scan() {
		folds := strings.Split(scanner2.Text(), "=")
		if folds[0] == "y" {
			horizontalSplit, _ := strconv.Atoi(folds[1])
			transparentIndex := 2
			for i := horizontalSplit + 1; i < y; i++ {
				for j := 0; j < x; j++ {
					points[i-transparentIndex][j] = points[i-transparentIndex][j] + points[i][j]
				}
				transparentIndex = transparentIndex + 2
			}
			y = horizontalSplit
		} else {
			verticalSplit, _ := strconv.Atoi(folds[1])
			transparentIndex := 2
			for i := verticalSplit + 1; i < x; i++ {
				for j := 0; j < y; j++ {
					points[j][i-transparentIndex] = points[j][i-transparentIndex] + points[j][i]
				}
				transparentIndex = transparentIndex + 2
			}
			x = verticalSplit
		}
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if points[i][j] > 0 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
