package wasteland

import (
	"fmt"
	trebuchet "github.com/matslabek/advent-of-code-23/1"
	"strings"
)

const FilePath = "8/input.txt"

func Wasteland() {
	inputStrings, err := trebuchet.ReadStringsFromFile(FilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	instructions := inputStrings[0]
	l := len(instructions)

	n := len(inputStrings)
	m := make(map[string][]string)

	for i := 2; i < n; i++ {
		s := strings.Fields(inputStrings[i])
		mm := make([]string, 0)
		mm = append(mm, s[2][1:4], s[3][0:3])
		m[s[0]] = mm
	}
	// Starting point should be AAA
	node := "AAA"

	counter := 1
	for i := 0; i < l; i++ {
		direction := instructions[i]
		v, _ := m[node]
		if direction == 'R' {
			node = v[1]
		} else if direction == 'L' {
			node = v[0]
		}

		if node == "ZZZ" {
			break
		}
		// If we reach the end of instruction strings go back to the beginning
		if i == l-1 {
			i = -1
		}
		counter++
	}
	fmt.Println(counter)
}
