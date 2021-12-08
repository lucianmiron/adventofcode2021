package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var uniqueNumbers [10]string = [10]string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}
	file, err := os.Open("C:\\Users\\miron\\Desktop\\AdventOfCode\\day8\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var oneScrambled string
	var fourScrambled string
	var sevenScrambled string

	var a, b, c, d, e, f, g string

	var sum int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		var line = scanner.Text()
		var lineComponents = strings.Split(line, "|")
		var length5ScrambledTogether string = ""
		var length6ScrambledTogether string = ""

		var output = lineComponents[1]
		var outputNumbers = strings.Split(output, " ")

		var uniqueNumberLine = strings.Split(lineComponents[0], " ")

		for i := 0; i < len(uniqueNumberLine); i++ {
			switch len(uniqueNumberLine[i]) {
			case 2:
				oneScrambled = uniqueNumberLine[i]
			case 3:
				sevenScrambled = uniqueNumberLine[i]
			case 4:
				fourScrambled = uniqueNumberLine[i]
			case 5:
				length5ScrambledTogether += uniqueNumberLine[i]
			case 6:
				length6ScrambledTogether += uniqueNumberLine[i]
			}
		}
		// A found
		for i := 0; i < len(sevenScrambled); i++ {
			if !strings.Contains(oneScrambled, string(sevenScrambled[i])) {
				a = string(sevenScrambled[i])
				break
			}
		}
		// CF found
		var cf = oneScrambled
		// BD Found
		var bd string
		for i := 0; i < len(fourScrambled); i++ {
			if !strings.Contains(oneScrambled, string(fourScrambled[i])) {
				bd += string(fourScrambled[i])
			}
		}
		// DG found
		var dg string
		for i := 0; i < len(length5ScrambledTogether); i++ {
			var occ int = 0
			for j := 0; j < len(length5ScrambledTogether); j++ {
				if length5ScrambledTogether[i] == length5ScrambledTogether[j] {
					occ++
				}
			}
			if occ == 3 {
				dg += string(length5ScrambledTogether[i])
				length5ScrambledTogether = strings.Replace(length5ScrambledTogether, string(length5ScrambledTogether[i]), " ", -1)
			}
		}
		dg = strings.Replace(dg, a, "", -1)
		// D found
		if bd[0] == dg[0] {
			d = string(bd[0])
		} else if bd[0] == dg[1] {
			d = string(bd[0])
		} else if bd[1] == dg[0] {
			d = string(bd[1])
		} else if bd[1] == dg[1] {
			d = string(bd[1])
		}
		// B found
		if string(bd[0]) == d {
			b = string(bd[1])
		} else {
			b = string(bd[0])
		}
		// G found
		if string(dg[0]) == d {
			g = string(dg[1])
		} else {
			g = string(dg[0])
		}
		// BFG found
		var bfg string
		for i := 0; i < len(length6ScrambledTogether); i++ {
			var occ int = 0
			for j := 0; j < len(length6ScrambledTogether); j++ {
				if length6ScrambledTogether[i] == length6ScrambledTogether[j] {
					occ++
				}
			}
			if occ == 3 {
				bfg += string(length6ScrambledTogether[i])
				length6ScrambledTogether = strings.Replace(length6ScrambledTogether, string(length6ScrambledTogether[i]), " ", -1)
			}
		}
		bfg = strings.Replace(bfg, a, "", -1)
		// F found
		f = strings.Replace(strings.Replace(bfg, b, "", -1), g, "", -1)
		// C found
		c = strings.Replace(cf, f, "", -1)
		// E found
		e = strings.Replace("abcdefg", a, "", -1)
		e = strings.Replace(e, b, "", -1)
		e = strings.Replace(e, c, "", -1)
		e = strings.Replace(e, d, "", -1)
		e = strings.Replace(e, f, "", -1)
		e = strings.Replace(e, g, "", -1)

		var realNumber int
		for i := 1; i < len(outputNumbers); i++ {
			var number string
			for j := 0; j < len(outputNumbers[i]); j++ {
				switch string(outputNumbers[i][j]) {
				case a:
					number += "a"
				case b:
					number += "b"
				case c:
					number += "c"
				case d:
					number += "d"
				case e:
					number += "e"
				case f:
					number += "f"
				case g:
					number += "g"
				}
			}
			for j := 0; j < len(uniqueNumbers); j++ {
				if len(string(uniqueNumbers[j])) != len(number) {
					continue
				}
				var found = true
				for k := 0; k < len(number); k++ {
					if !strings.Contains(uniqueNumbers[j], string(number[k])) {
						found = false
					}
				}
				if found {
					realNumber = realNumber + j
					realNumber = realNumber * 10
				}
			}
		}
		sum += realNumber / 10
	}

	fmt.Printf("%d", sum)
}
