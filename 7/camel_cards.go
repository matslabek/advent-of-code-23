package camel

import (
	trebuchet "AdventOfCode/1"
	"fmt"
)

const FilePath = "7/input.txt"

func CamelCards() {
	inputStrings, err := trebuchet.ReadStringsFromFile(FilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
