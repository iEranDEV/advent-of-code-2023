package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	winningNumbers []int
	yourNumbers    []int
}

func main() {
	file, err := os.ReadFile("input.txt")
	content := strings.ReplaceAll(string(file), "\r", "")
	content = strings.ReplaceAll(content, "Card ", "")
	if err != nil {
		log.Fatal(err)
	}

	cards := []card{}
	for _, row := range strings.Split(content, "\n") {
		data := strings.Split(row[strings.Index(row, ":")+2:], " | ")
		winningNumbers := []int{}
		yourNumbers := []int{}
		for _, value := range strings.Split(data[0], " ") {
			number, err := strconv.Atoi(value)
			if err == nil {
				winningNumbers = append(winningNumbers, number)
			}
		}
		for _, value := range strings.Split(data[1], " ") {
			number, err := strconv.Atoi(value)
			if err == nil {
				yourNumbers = append(yourNumbers, number)
			}
		}
		cards = append(cards, card{winningNumbers: winningNumbers, yourNumbers: yourNumbers})
	}

	// PART 1
	points := 0
	for _, card := range cards {
		currPoints := 0
		for _, yourNumber := range card.yourNumbers {
			if slices.Contains(card.winningNumbers, yourNumber) {
				if currPoints == 0 {
					currPoints = 1
				} else {
					currPoints *= 2
				}
			}
		}
		points += currPoints
	}
	fmt.Println("Part 1:", points)

	// PART 2
	amounts := make([]int, len(cards))
	for index, card := range cards {
		currPoints := 0
		for _, yourNumber := range card.yourNumbers {
			if slices.Contains(card.winningNumbers, yourNumber) {
				currPoints++
			}
		}
		amount := amounts[index]
		for i := 0; i < currPoints; i++ {
			amounts[index+i+1] += amount + 1
		}
	}
	sum := 0
	for _, val := range amounts {
		sum += val + 1
	}
	fmt.Println("Part 2:", sum)
}
