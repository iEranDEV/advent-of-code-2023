package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func part1(content string) {
	count := 0
	for _, val := range strings.Split(content, "\n") {
		result := []int{}
		for _, letter := range val {
			if unicode.IsNumber(letter) {
				n, _ := strconv.Atoi(string(letter))
				result = append(result, n)
			}
		}
		if len(result) > 0 {
			count += result[len(result)-1] + (10 * (result[0]))
		}
	}
	fmt.Println(count)
}

func part2(content string) {
	count := 0
	digits := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, val := range strings.Split(content, "\n") {
		result := []int{}
		for index, letter := range val {
			if unicode.IsNumber(letter) {
				n, _ := strconv.Atoi(string(letter))
				result = append(result, n)
			} else {
				for i, digit := range digits {
					if index+len(digit) <= len(val) && val[index:index+len(digit)] == digit {
						result = append(result, i+1)
					}
				}
			}
		}
		if len(result) > 0 {
			count += result[len(result)-1] + (10 * (result[0]))
			fmt.Println(result)
		}
	}
	fmt.Println(count)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// part1(string(content))
	part2(string(content))
}
