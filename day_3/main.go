package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type gear struct {
	y int
	x int
}

var (
	schema []string
	width  int
	height int
	gears  = make(map[gear][]int)
)

func checkRelatives(i int, j int, currNumber string, sum *int) {
	nearSymbols := false
	for y := i - 1; y <= i+1; y++ {
		for x := j - len(currNumber) - 1; x <= j; x++ {
			if y >= 0 && y < height && x >= 0 && x < width {
				if !unicode.IsNumber(rune(schema[y][x])) && string(schema[y][x]) != string('.') {
					nearSymbols = true
					if string(schema[y][x]) == "*" {
						temp := gear{y: y, x: x}
						val, ok := gears[temp]
						number, err := strconv.Atoi(currNumber)
						if ok && err == nil {
							gears[temp] = append(val, number)
						} else {
							gears[temp] = []int{number}
						}
					}
				}
			}
		}
	}
	if nearSymbols {
		number, err := strconv.Atoi(currNumber)
		if err == nil {
			*sum += number
		}
	}
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	schema = strings.Split(strings.ReplaceAll(string(content), "\r", ""), "\n")
	width = len(schema[0])
	height = len(schema)

	// Part 1
	sum := 0
	for i, row := range schema {
		currNumber := ""
		for j, letter := range strings.Split(row, "") {
			val, err := strconv.Atoi(letter)
			if err == nil {
				currNumber += strconv.Itoa(val)
			} else if currNumber != "" {
				checkRelatives(i, j, currNumber, &sum)
				currNumber = ""
			}
		}
		checkRelatives(i, width, currNumber, &sum)
	}
	fmt.Println("Part 1:", sum)

	// Part 2
	gearRatio := 0
	for _, val := range gears {
		if len(val) == 2 {
			gearRatio += val[0] * val[1]
		}
	}
	fmt.Println("Part 2:", gearRatio)
}
