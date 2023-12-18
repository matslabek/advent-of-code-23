package camel

import (
	trebuchet "AdventOfCode/1"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const FilePath = "7/input.txt"

func CamelCards() {

	// We'll use cardTypes to group hands by types, in descending order, so five of a kind, four of a kind, full house etc
	cardTypes := make([][]string, 0)

	fiveOfKind := make([]string, 0)
	fourOfKind := make([]string, 0)
	fullHouse := make([]string, 0)
	threeOfKind := make([]string, 0)
	twoPair := make([]string, 0)
	onePair := make([]string, 0)
	highCard := make([]string, 0)

	inputStrings, err := trebuchet.ReadStringsFromFile(FilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	cardBids := make(map[string]int, len(inputStrings))
	for _, inpstr := range inputStrings {
		splitRes := strings.Fields(inpstr)
		bid, _ := strconv.Atoi(splitRes[1])
		cardBids[splitRes[0]] = bid
		switch handEvaluator(splitRes[0]) {
		case 6:
			fiveOfKind = append(fiveOfKind, splitRes[0])
			break
		case 5:
			fourOfKind = append(fourOfKind, splitRes[0])
			break
		case 4:
			fullHouse = append(fullHouse, splitRes[0])
			break
		case 3:
			threeOfKind = append(threeOfKind, splitRes[0])
			break
		case 2:
			twoPair = append(twoPair, splitRes[0])
			break
		case 1:
			onePair = append(onePair, splitRes[0])
			break
		case 0:
			highCard = append(highCard, splitRes[0])
			break
		}
	}
	cardTypes = append(cardTypes, fiveOfKind, fourOfKind, fullHouse, threeOfKind, twoPair, onePair, highCard)
	// Base rank - calculated base on what's the card type
	baseValByType := make([]int, 7)
	baseVal := 0
	for i := 6; i >= 0; i-- {
		if i != 6 {
			baseVal += len(cardTypes[i+1])
		}
		baseValByType[i] = baseVal
	}
	totalBid := 0
	for typeNumber, sl := range cardTypes {
		sort.Slice(sl, func(i, j int) bool {
			return singleCardComparer(sl[i], sl[j])
		})
		for i := len(sl) - 1; i >= 0; i-- {
			rank := (baseValByType[typeNumber] + i + 1)
			totalBid += rank * cardBids[sl[i]]
		}
	}
	fmt.Println(totalBid)
}

func handEvaluator(hand string) int {
	cardMap := make(map[string]int)
	for i := 0; i < 5; i++ {
		_, ok := cardMap[string(hand[i])]
		if ok {
			cardMap[string(hand[i])]++
		} else {
			cardMap[string(hand[i])] = 1
		}
	}

	cardMapValues := make([]int, 0)
	for _, v := range cardMap {
		cardMapValues = append(cardMapValues, v)
	}

	eval := 1
	switch len(cardMapValues) {
	// Five of a kind
	case 1:
		return 6
	case 2:
		// Four of a kind
		if cardMapValues[0] == 4 || cardMapValues[0] == 1 {
			return 5
		} else {
			// Fullhouse
			return 4
		}
	case 3:
		if cardMapValues[0] == 3 {
			// Three of a kind
			return 3
		} else if cardMapValues[0] == 2 {
			// Two pair
			return 2
		} else {
			if cardMapValues[1] == 1 || cardMapValues[1] == 3 {
				return 3
			} else {
				return 2
			}
		}
	case 4:
		// One pair
		return 1
	case 5:
		// High card
		return 0
	}

	for _, v := range cardMap {
		if v > eval {
			eval = v
		}
	}
	return eval
}

// Returns true if hand1 < hand2
func singleCardComparer(hand1, hand2 string) bool {
	cardValues := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": 11,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	for i := 0; i < 5; i++ {
		if cardValues[string(hand1[i])] > cardValues[string(hand2[i])] {
			return false
		} else if cardValues[string(hand1[i])] < cardValues[string(hand2[i])] {
			return true
		}
	}
	// This should never happen, cause hands are unique
	return false
}
