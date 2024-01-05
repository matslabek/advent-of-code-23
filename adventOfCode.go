package main

import (
	"fmt"
	trebuchet "github.com/matslabek/advent-of-code-23/1"
	puzzle "github.com/matslabek/advent-of-code-23/9"
)

func main() {
	input, _ := trebuchet.ReadStringsFromFile("9/input.txt")
	fmt.Println(puzzle.Oasis(input))
}
