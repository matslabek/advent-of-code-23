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
	points := 0
	for _, inputString := range inputStrings {
		matches := 0
		splitCard := strings.Split(inputString, ":")
		splitNumbers := strings.Split(splitCard[1], "|")
		winningNumbers := strings.Fields(splitNumbers[0])
		myNumbers := strings.Fields(splitNumbers[1])
		for _, nr := range myNumbers {
			if contains(winningNumbers, nr) {
				matches++
			}
		}
		gamePoints := 0
		for i := 0; i < matches; i++ {
			if i == 0 {
				gamePoints = 1
			} else {
				gamePoints = gamePoints * 2
			}
		}
		points += gamePoints
	}
	fmt.Println(points)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
