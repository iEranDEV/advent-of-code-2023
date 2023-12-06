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

func RemoveIndex(s []numbersRange, index int) []numbersRange {
	return append(s[:index], s[index+1:]...)
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

	ranges := []numbersRange{}
	for i := 0; i < len(seedText); i += 2 {
		start, _ := strconv.Atoi(seedText[i])
		end, _ := strconv.Atoi(seedText[i+1])
		ranges = append(ranges, numbersRange{start: start, end: start + end - 1})
	}

	// For each mapping
	for _, operationsList := range operations {
		temp := []numbersRange{}
		for _, value := range ranges {
			replace := []numbersRange{}

			for _, operation := range operationsList {
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
				} else if operation.source.start >= value.start && operation.source.end > value.end && operation.source.start <= value.end {
					// NEW SEGMENT IS INSIDE FROM LEFT AND OUTSIDE FROM RIGHT

					diff := operation.destination.start - operation.source.start
					if operation.source.start == value.end {
						replace = append(replace, numbersRange{value.start, value.end - 1})
						replace = append(replace, numbersRange{value.end + diff, value.end + diff})
					} else {
						if operation.source.start == value.start {
							replace = append(replace, numbersRange{value.start + diff, value.end + diff})
						} else {
							replace = append(replace, numbersRange{value.start, operation.source.start - 1})
							replace = append(replace, numbersRange{operation.destination.start, value.end + diff})
						}

					}

				} else if operation.source.start < value.start && operation.source.end <= value.end && operation.source.end >= value.start {
					// NEW SEGMENT IS INSIDE FROM RIGHT AND OUTSIDE FROM LEFT

					diff := operation.destination.start - operation.source.start
					if operation.source.end == value.start {
						replace = append(replace, numbersRange{value.start + diff, value.start + diff})
						replace = append(replace, numbersRange{value.start + 1, value.end})
					} else {
						if operation.source.end == value.end {
							replace = append(replace, numbersRange{value.start + diff, value.end + diff})
						} else {
							replace = append(replace, numbersRange{value.start + diff, operation.destination.end})
							replace = append(replace, numbersRange{operation.source.end + 1, value.end})
						}
					}

				} else if operation.source.start < value.start && operation.source.end > value.end {
					// NEW SEGMENT IS FULLY OUTSIDE

					diff := operation.destination.start - operation.source.start
					replace = append(replace, numbersRange{value.start + diff, value.end + diff})
				}
			}
			if len(replace) >= 1 {
				temp = append(temp, replace...)
			} else {
				temp = append(temp, value)
			}
		}
		ranges = append([]numbersRange{}, temp...)
	}
	min := ranges[len(ranges)-1].start
	for _, value := range ranges {
		if value.start < min {
			min = value.start
		}
	}
	fmt.Println("Part 2:", min)
}

func main() {
	part2()
}
