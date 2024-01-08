package main

import (
	"fmt"
	trebuchet "github.com/matslabek/advent-of-code-23/1"
	puzzle "github.com/matslabek/advent-of-code-23/10"
)

func main() {
	input, _ := trebuchet.ReadStringsFromFile("10/input.txt")
	fmt.Println(puzzle.PathFinder(input))
}
