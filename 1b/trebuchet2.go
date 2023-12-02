package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Read input strings
func readStringsFromFile(filePath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var strings []string

	// Read each line and append it to the strings slice
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return strings, nil
}

func elvishCount(word string) int {

	// Workaround for cases like "oneight" with overlapping chars
	word = duplicateCharacters(word)

	// Split the string
	re := regexp.MustCompile(`\d|(one|two|three|four|five|six|seven|eight|nine)`)
	matches := re.FindAllString(word, -1)

	// Edge case - only one digit
	if len(matches) == 1 {
		matches = append(matches, matches[0])
	}

	firstDigit := matches[0]
	lastDigit := matches[len(matches)-1]

	// Matched "digit" is a string
	if len(firstDigit) > 1 {
		firstDigit = stringToDigit(firstDigit)
	}
	if len(lastDigit) > 1 {
		lastDigit = stringToDigit(lastDigit)
	}

	finalNumber, _ := strconv.Atoi(firstDigit + lastDigit)
	return finalNumber
}

func main() {
	filePath := "1/input.txt"

	strings, err := readStringsFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	sum := 0
	// Count the sum
	for _, str := range strings {
		value := elvishCount(str)
		sum += value
	}
	fmt.Print(sum)
}

func stringToDigit(str string) string {
	dM := make(map[string]string)
	dM["one"] = "1"
	dM["two"] = "2"
	dM["three"] = "3"
	dM["four"] = "4"
	dM["five"] = "5"
	dM["six"] = "6"
	dM["seven"] = "7"
	dM["eight"] = "8"
	dM["nine"] = "9"

	value, ok := dM[str]
	if ok {
		return value
	}
	return "0"
}

// Obnoxious workaround because go re2 has no positive overlook, so we need to prepare the string for the regex
func duplicateCharacters(input string) string {
	s1 := strings.SplitAfter(input, "one")
	i1 := strings.Join(s1, "e")
	s2 := strings.SplitAfter(i1, "two")
	i2 := strings.Join(s2, "o")
	s3 := strings.SplitAfter(i2, "three")
	i3 := strings.Join(s3, "e")
	s5 := strings.SplitAfter(i3, "five")
	i5 := strings.Join(s5, "e")
	s7 := strings.SplitAfter(i5, "seven")
	i7 := strings.Join(s7, "n")
	s8 := strings.SplitAfter(i7, "eight")
	i8 := strings.Join(s8, "t")
	s9 := strings.SplitAfter(i8, "nine")
	i9 := strings.Join(s9, "e")
	return i9
}
