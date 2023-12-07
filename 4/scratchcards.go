package scratchcards

import (
	treb "AdventOfCode/1"
	"fmt"
	"strings"
)

func Scratch() {
	inputStrings, err := treb.ReadStringsFromFile("4/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// Points for part one
	points := 0

	nrOfCards := len(inputStrings)

	cardsMap := make(map[int]int)
	for i := 1; i <= nrOfCards; i++ {
		cardsMap[i] = 1
	}
	for cardNumber, inputString := range inputStrings {
		cardNumber += 1 // Start counting from 1
		matches := 0
		multiplier := cardsMap[cardNumber]
		splitCard := strings.Split(inputString, ":")
		splitNumbers := strings.Split(splitCard[1], "|")
		winningNumbers := strings.Fields(splitNumbers[0])
		myNumbers := strings.Fields(splitNumbers[1])
		for _, nr := range myNumbers {
			if contains(winningNumbers, nr) {
				matches++
			}
		}

		// Only relevant to part 1:
		gamePoints := 0
		for i := 0; i < matches; i++ {
			if i == 0 {
				gamePoints = 1
			} else {
				gamePoints = gamePoints * 2
			}
		}
		points += gamePoints
		// ---

		// Part 2 of the puzzle
		for i := 0; i < matches; i++ {
			cardsMap[cardNumber+1+i] += multiplier
		}
	}
	//Part 1 answer
	fmt.Println(points)

	sum := 0
	for i := 1; i <= nrOfCards; i++ {
		sum += cardsMap[i]
	}
	// Part 2 answer
	fmt.Println(sum)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
