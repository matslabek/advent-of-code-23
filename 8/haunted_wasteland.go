package wasteland

import (
	"fmt"
	trebuchet "github.com/matslabek/advent-of-code-23/1"
	"strings"
)

const FilePath = "8/input.txt"

// Least common multiple
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

// Greatest Common Divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Part1() {
	inputStrings, _ := trebuchet.ReadStringsFromFile(FilePath)

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

// Part two solution depends on the nodes cycle, and uses maths LCM and GCD
func Part2() {
	inputStrings, err := trebuchet.ReadStringsFromFile(FilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	instructions := inputStrings[0]
	l := len(instructions)

	n := len(inputStrings)
	m := make(map[string][]string)

	startingNodes := make([]string, 0)

	for i := 2; i < n; i++ {
		s := strings.Fields(inputStrings[i])
		mm := make([]string, 0)
		mm = append(mm, s[2][1:4], s[3][0:3])
		m[s[0]] = mm
		if s[0][2] == 'A' {
			startingNodes = append(startingNodes, s[0])
		}
	}
	// Starting nodes are all nodes that end with 'A'

	// Calculate the path for each starting node:
	cycles := make([]int, 0)

	lj := len(startingNodes)
	for j := 0; j < lj; j++ {
		counter := 1
		for i := 0; i < l; i++ {
			direction := instructions[i]
			node := startingNodes[j]
			v, _ := m[node]
			if direction == 'R' {
				startingNodes[j] = v[1]
			} else if direction == 'L' {
				startingNodes[j] = v[0]
			}
			if startingNodes[j][2] == 'Z' {
				cycles = append(cycles, counter)
				break
			}
			counter++
			// If we reach the end of instruction strings go back to the beginning
			if i == l-1 {
				i = -1
			}
		}
	}

	val := cycles[0]
	for i := 1; i < len(cycles); i++ {
		val = lcm(val, cycles[i])
	}

	fmt.Println(val)
}
