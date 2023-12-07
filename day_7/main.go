package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards []int
	bid   int
}

type HandSort []hand

func (a HandSort) Len() int {
	return len(a)
}

func getHandType(cards []int) int {
	data := map[int]int{}
	for i := 0; i < 5; i++ {
		val, ok := data[cards[i]]
		if ok {
			data[cards[i]] = val + 1
		} else {
			data[cards[i]] = 1
		}
	}
	result := 0
	values := []int{}
	for _, v := range data {
		values = append(values, v)
	}
	if len(data) == 1 {
		result = 7
	} else if len(data) == 2 && (values[0] == 4 || values[1] == 4) {
		result = 6
	} else if len(data) == 2 && (values[0] == 3 || values[1] == 3) {
		result = 5
	} else if len(data) == 3 && (values[0] == 3 || values[1] == 3 || values[2] == 3) {
		result = 4
	} else if len(data) == 3 {
		result = 3
	} else if len(data) == 4 {
		result = 2
	} else {
		result = 1
	}
	return result
}

func (a HandSort) Less(i, j int) bool {
	isLess := getHandType(a[i].cards) < getHandType(a[j].cards)
	if isLess {
		return true
	} else if getHandType(a[i].cards) > getHandType(a[j].cards) {
		return false
	} else if getHandType(a[i].cards) == getHandType(a[j].cards) {
		for x := 0; x < 5; x++ {
			if a[j].cards[x] != a[i].cards[x] {
				if a[j].cards[x] < a[i].cards[x] {
					return false
				} else {
					return true
				}
			}
		}
		return false
	}
	return false
}

func (a HandSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

var cardPoints []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

func main() {
	file, err := os.ReadFile("input.txt")
	content := strings.Split(strings.ReplaceAll(string(file), "\r", ""), "\n")
	if err != nil {
		log.Fatal(err)
	}
	hands := []hand{}
	for _, line := range content {
		val := strings.Split(line, " ")
		data := []int{}
		for _, card := range val[0] {
			data = append(data, slices.Index(cardPoints, string(card)))
		}
		bid, _ := strconv.Atoi(val[1])
		hands = append(hands, hand{cards: data, bid: bid})
	}
	sort.Sort(HandSort(hands))
	result := 0
	for i := 1; i <= len(hands); i++ {
		result += i * hands[i-1].bid
	}
	fmt.Println("Part 1:", result)
}
