package trebuchet

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Read input strings
func ReadStringsFromFile(filePath string) ([]string, error) {
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
	// Split the string
	re := regexp.MustCompile(`\d`)
	matches := re.FindAllString(word, -1)

	// Edge case - only one digit
	if len(matches) == 1 {
		matches = append(matches, matches[0])
	}

	firstDigit := matches[0]
	lastDigit := matches[len(matches)-1]

	finalNumber, _ := strconv.Atoi(firstDigit + lastDigit)
	return finalNumber
}

func Trebuchet() {
	filePath := "1/input.txt"

	strings, err := ReadStringsFromFile(filePath)
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
