package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Part 1
func part1() {
	file, err := os.ReadFile("input.txt")
	content := strings.ReplaceAll(string(file), "\r", "")
	if err != nil {
		log.Fatal(err)
	}
	splitContent := strings.Split(content, "\n")

	// Initialize seeds list
	seedsData := [][]int{}
	for _, val := range strings.Split(strings.ReplaceAll(splitContent[0], "seeds: ", ""), " ") {
		number, err := strconv.Atoi(val)
		if err == nil {
			seedsData = append(seedsData, []int{number})
		}
	}

	index := 1
	skip := true
	for _, row := range splitContent[2:] {
		if row == "" {
			index++
			for j, seed := range seedsData {
				if len(seed) < index {
					seedsData[j] = append(seedsData[j], seed[len(seed)-1])
				}
			}
			skip = true
		} else if skip {
			skip = false
			continue
		} else {
			data := []int{}
			for _, val := range strings.Split(row, " ") {
				number, err := strconv.Atoi(val)
				if err == nil {
					data = append(data, number)
				}
			}
			for j, seed := range seedsData {
				if seed[index-1] >= data[1] && seed[index-1] < data[1]+data[2] {
					for i := 0; i < data[2]; i++ {
						sourceNumber := data[1] + i
						destinationNumber := data[0] + i
						if seed[index-1] == sourceNumber {
							seedsData[j] = append(seedsData[j], destinationNumber)
						}
					}
				}
			}
		}
	}
	minLocation := seedsData[0][len(seedsData[0])-1]
	for _, seed := range seedsData {
		if seed[len(seed)-1] < minLocation {
			minLocation = seed[len(seed)-1]
		}
	}
	fmt.Println("Part 1:", minLocation)
}

type numbersRange struct {
	start int
	end   int
}

type operation struct {
	destionation int
	source       int
	amount       int
}

// Part 2
func part2() {
	file, err := os.ReadFile("input.txt")
	content := strings.ReplaceAll(string(file), "\r", "")
	if err != nil {
		log.Fatal(err)
	}
	splitContent := strings.Split(content, "\n")
	seedText := strings.Split(strings.ReplaceAll(splitContent[0], "seeds: ", ""), " ")

	ranges := [][]numbersRange{{}}
	for i := 0; i < len(seedText); i += 2 {
		start, _ := strconv.Atoi(seedText[i])
		end, _ := strconv.Atoi(seedText[i+1])
		ranges[0] = append(ranges[0], numbersRange{start: start, end: start + end - 1})
	}

	operations := [][]operation{}
	index := -1
	skip := true
	for _, row := range splitContent[2:] {
		if row == "" {
			skip = true
		} else if skip {
			skip = false
			operations = append(operations, []operation{})
			index++
			continue
		} else {
			data := []int{}
			for _, val := range strings.Split(row, " ") {
				number, err := strconv.Atoi(val)
				if err == nil {
					data = append(data, number)
				}
			}
			operations[index] = append(operations[index], operation{destionation: data[0], source: data[1], amount: data[2]})
		}
	}
	for _, i := range operations {
		fmt.Println(i)
	}
	/*
		rangeToSwap := numbersRange{data[1], data[1] + data[2] - 1}
				newRange := numbersRange{data[0], data[0] + data[2] - 1}
				temp := []numbersRange{}
				for _, value := range ranges[index] {
					replace := []numbersRange{}
					if rangeToSwap.start >= value.start {
						replace = append(replace, numbersRange{start: newRange.start, end: newRange.end})

						if rangeToSwap.start != value.start {
							replace = append(replace, numbersRange{start: value.start, end: rangeToSwap.start - 1})
						}

						if rangeToSwap.end != value.end {
							replace = append(replace, numbersRange{rangeToSwap.end + 1, value.end})
						}
					}
					if len(replace) >= 1 {
						temp = append(temp, replace...)
					} else {
						temp = append(temp, value)
					}
				}
				fmt.Println("temp", temp)
	*/
	// for _, val := range ranges {
	// 	fmt.Println(val)
	// }
	// minLocation := seedsData[0][len(seedsData[0])-1]
	// for _, seed := range seedsData {
	// 	if seed[len(seed)-1] < minLocation {
	// 		minLocation = seed[len(seed)-1]
	// 	}
	// }
	// fmt.Println("Part 1:", minLocation)
}

func main() {
	part1()
}
