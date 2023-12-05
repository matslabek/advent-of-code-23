package gear_ratios

import (
	trebuchet "AdventOfCode/1"
	"fmt"
	"strconv"
)

const FilePath = "3/input.txt"
const RowLength = 140
const ColumnHeight = 140

var symbols = [10]string{"&", "*", "/", "$", "@", "=", "-", "#", "+", "%"}

func GearRatios() {
	inputStrings, err := trebuchet.ReadStringsFromFile(FilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	sum := 0
	for x, str := range inputStrings {
		insideNumberString := false
		numberString := ""
		var numberStart *int
		var numberEnd *int
		for y, char := range str {
			// Char is a digit
			if _, err := strconv.Atoi(string(char)); err == nil {
				numberString += string(char)
				// It's the first digit of the number
				if !insideNumberString {
					insideNumberString = true
					start := y
					numberStart = &start
				}
				// It's the end of line
				if y == RowLength-1 {
					insideNumberString = false
					numberEnd = &y
				} else {
					// It's the end of the numberString
					if _, err := strconv.Atoi(string(str[y+1])); err != nil {
						insideNumberString = false
						end := y
						numberEnd = &end
					}
				}
				// We got the number, check if it neighbors a symbol
				if numberEnd != nil {
					if checkForSymbolsNearby(x, *numberStart, *numberEnd, inputStrings) {
						fmt.Println(numberString)
						value, _ := strconv.Atoi(numberString)
						sum += value
					}
					numberStart = nil
					numberEnd = nil
					numberString = ""
				}
			}
		}
	}
	fmt.Println(sum)
}
func checkForSymbolsNearby(rowNumber, start, end int, rowsAndColumns []string) bool {
	numberLength := end - start + 1
	//Check left
	if start != 0 && contains(symbols, string(rowsAndColumns[rowNumber][start-1])) {
		return true
	}
	//Check top left
	if start != 0 && rowNumber != 0 && contains(symbols, string(rowsAndColumns[rowNumber-1][start-1])) {
		return true
	}
	//Check bottom left
	if start != 0 && rowNumber != ColumnHeight-1 && contains(symbols, string(rowsAndColumns[rowNumber+1][start-1])) {
		return true
	}
	//Check right
	if end != RowLength-1 && contains(symbols, string(rowsAndColumns[rowNumber][end+1])) {
		return true
	}
	//Check top right
	if end != RowLength-1 && rowNumber != 0 && contains(symbols, string(rowsAndColumns[rowNumber-1][end+1])) {
		return true
	}
	//Check bottom right
	if end != RowLength-1 && rowNumber != ColumnHeight-1 && contains(symbols, string(rowsAndColumns[rowNumber+1][end+1])) {
		return true
	}

	for i := 0; i < numberLength; i++ {
		//Check up
		if rowNumber != 0 && contains(symbols, string(rowsAndColumns[rowNumber-1][start+i])) {
			return true
		}
		//Check down
		if rowNumber != ColumnHeight-1 && contains(symbols, string(rowsAndColumns[rowNumber+1][start+i])) {
			return true
		}
	}
	return false
}

func contains(s [10]string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
