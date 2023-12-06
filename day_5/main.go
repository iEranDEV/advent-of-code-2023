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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type operation struct {
	destination numbersRange
	source      numbersRange
	amount      int
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
			operations[index] = append(operations[index], operation{destination: numbersRange{data[0], data[0] + data[2] - 1}, source: numbersRange{data[1], data[1] + data[2] - 1}, amount: data[2]})
		}
	}

	ranges := [][]numbersRange{}
	for i := 0; i < len(seedText); i += 2 {
		start, _ := strconv.Atoi(seedText[i])
		end, _ := strconv.Atoi(seedText[i+1])
		ranges = append(ranges, []numbersRange{})
		ranges[0] = append(ranges[0], numbersRange{start: start, end: start + end - 1})
	}

	index = 0
	for _, operationsList := range operations {
		temp := []numbersRange{}
		for _, operation := range operationsList {
			for _, value := range ranges[index] {
				replace := []numbersRange{}

				if operation.source.start >= value.start && operation.source.end <= value.end {
					// NEW SEGMENT IS FULL IN OLD ONE

					if operation.source.start != value.start {
						if operation.source.end != value.end {
							replace = append(replace, numbersRange{value.start, operation.source.start - 1})
							replace = append(replace, numbersRange{operation.destination.start, operation.destination.end})
							replace = append(replace, numbersRange{operation.source.end + 1, value.end})
						} else {
							replace = append(replace, numbersRange{value.start, operation.source.start - 1})
							replace = append(replace, numbersRange{operation.destination.start, operation.destination.end})
						}
					} else {
						if operation.source.end != value.end {
							replace = append(replace, numbersRange{operation.destination.start, operation.destination.end})
							replace = append(replace, numbersRange{operation.source.end + 1, value.end})
						} else {
							replace = append(replace, numbersRange{operation.destination.start, operation.destination.end})
						}
					}
				} else if operation.source.start >= value.start && operation.source.end > value.end {
					// NEW SEGMENT IS INSIDE FROM LEFT AND OUTSIDE FROM RIGHT

				} else if operation.source.start < value.start && operation.source.end <= value.end {
					// NEW SEGMENT IS INSIDE FROM RIGHT AND OUTSIDE FROM LEFT

				} else if operation.source.start < value.start && operation.source.end > value.end {
					// NEW SEGMENT IS FULLY OUTSIDE

					replace = append(replace, numbersRange{value.start + operation.amount, value.end + operation.amount})
				}

				if len(replace) >= 1 {
					temp = append(temp, replace...)
				} else {
					temp = append(temp, value)
				}
			}
		}
		if len(ranges)-1 <= index {
			ranges = append(ranges, []numbersRange{})
		}
		ranges[index+1] = append(ranges[index+1], temp...)
		index++
	}
}

func main() {
	part2()
}
