package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day14\\input.txt")

	insertionRules := ""
	polymerTemplate := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if polymerTemplate == "" {
			polymerTemplate = scanner.Text()
		} else {
			insertionRules += scanner.Text() + "\n"
		}
	}
	noSteps := 40
	for i := 0; i < noSteps; i++ {
		nextPolymer := ""
		for j := 0; j < len(polymerTemplate)-1; j++ {
			ruleIndex := strings.Index(insertionRules, string(polymerTemplate[j])+string(polymerTemplate[j+1]))
			if ruleIndex != -1 {
				nextPolymer += string(insertionRules[ruleIndex]) + string(insertionRules[ruleIndex+6])
			}
		}
		polymerTemplate = nextPolymer + string(polymerTemplate[len(polymerTemplate)-1])
		fmt.Printf("Step:%d", i)
	}

	// var occurrences [25]int
	// for i := 0; i < len(polymerTemplate); i++ {
	// 	occurrences[polymerTemplate[i]-65]++
	// }

	// max := 0
	// min := int(^uint(0) >> 1)
	// for i := 0; i < len(occurrences); i++ {
	// 	if occurrences[i] == 0 {
	// 		continue
	// 	}
	// 	if occurrences[i] > max {
	// 		max = occurrences[i]
	// 	}
	// 	if occurrences[i] < min {
	// 		min = occurrences[i]
	// 	}
	// }
	// delta := max - min
	// fmt.Printf("%d", delta)
}
