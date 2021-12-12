package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day10\\input-easy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	openingChunks := "([{<"
	// parantezeRotunde := "()"
	// parantezeDrepte := "[]"
	// acolade := "{}"
	// micmare := "<>"
	points := 0
	currentPoint :=0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		lineC := []rune(line)
		for i := 1; i < len(lineC); i++ {
			symbol := string(lineC[i])
			prevSymbol := string(lineC[i-1])
			if !strings.Contains(openingChunks, string(lineC[i])) {
				// Enclosing Chunk
				if symbol == ")" && prevSymbol != "(" {
					points += 3
					break
				} else if symbol == "]" && prevSymbol != "[" {
					points += 57
					break
				} else if symbol == "}" && prevSymbol != "{" {
					points += 1197
					break
				} else if symbol == ">" && prevSymbol != "<" {
					points += 25137
					break
				}
				// Reduce echunks
				// line = strings.Replace(line, string(line[i-1]), "", 1)
				// line = strings.Replace(line, string(line[i]), "", 1)
				lineC = append(lineC[:i], lineC[i+2:]...)
				i = i - 1
			}
		}
	}
	fmt.Printf("%d", points)
}
