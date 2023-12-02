package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var params = map[string]int{"red": 12, "green": 13, "blue": 14}

func part1(content string) {
	result := 0
	for i, text := range strings.Split(content, "\n") {
		sets := strings.Split(strings.Split(text, ": ")[1], "; ")
		possible := true
		for _, set := range sets {
			for _, data := range strings.Split(set, ", ") {
				splitData := strings.Split(data, " ")
				a := params[strings.ReplaceAll(splitData[1], "\r", "")]
				n, _ := strconv.Atoi(splitData[0])
				if a < n {
					possible = false
				}
			}
		}
		if possible {
			result += i + 1
		}
	}
	fmt.Println(result)
}

func part2(content string) {
	sum := 0
	for _, text := range strings.Split(content, "\n") {
		sets := strings.Split(strings.Split(text, ": ")[1], "; ")
		values := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, set := range sets {
			for _, data := range strings.Split(set, ", ") {
				splitData := strings.Split(data, " ")
				color := strings.ReplaceAll(splitData[1], "\r", "") // Color
				amountOfCubes, _ := strconv.Atoi(splitData[0])      // Amount of cubes
				if amountOfCubes > values[color] {
					values[color] = amountOfCubes
				}
			}
		}
		result := 1
		for _, val := range values {
			result *= val
		}
		sum += result
	}
	fmt.Println(sum)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// part1(string(content))
	part2(string(content))
}
