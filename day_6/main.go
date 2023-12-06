package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type race struct {
	time   int
	record int
}

func main() {
	file, err := os.ReadFile("input.txt")
	content := strings.Split(strings.ReplaceAll(string(file), "\r", ""), "\n")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	data := []race{}
	result := 1
	space := regexp.MustCompile(`\s+`)
	times := strings.Split(space.ReplaceAllString(content[0], " "), " ")[1:]
	distances := strings.Split(space.ReplaceAllString(content[1], " "), " ")[1:]
	size := len(times)
	for i := 0; i < size; i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		data = append(data, race{time, distance})
	}
	for _, value := range data {
		possible := 0
		for i := 0; i < value.time; i++ {
			// i is equal to speed/milisecond
			restTime := value.time - i
			distance := restTime * i
			if distance > value.record {
				possible += 1
			}
		}
		if possible > 0 {
			result *= possible
		}
	}
	fmt.Println("Part 1:", result)

	// Part 2
	finalRaceTime, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(content[0], " ", ""), "Time:", ""))
	finalRaceDistance, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(content[1], " ", ""), "Distance:", ""))
	possible := 0
	for i := 0; i < finalRaceTime; i++ {
		restTime := finalRaceTime - i
		distance := restTime * i
		if distance > finalRaceDistance {
			possible += 1
		}
	}
	fmt.Println("Part 2:", possible)
}
