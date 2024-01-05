package mirage

import (
	"strconv"
	"strings"
)

func Part1(inputs []string) int {
	li := len(inputs)
	result := 0
	for i := 0; i < li; i++ {
		history := parseInputData(inputs[i])
		result += solve(history)
	}
	return result
}

// The task begs to be solved recursively
func solve(dataset []int) int {
	sum := 0
	for _, v := range dataset {
		sum += v
	}
	if sum == 0 {
		return 0
	} else {
		l := len(dataset)
		childDataset := make([]int, l-1)
		for i := 0; i < l-1; i++ {
			childDataset[i] = dataset[i+1] - dataset[i]
		}
		return solve(childDataset) + dataset[l-1]
	}
}
func parseInputData(input string) []int {
	d := strings.Fields(input)
	dataset := make([]int, len(d))
	for i, s := range d {
		s, _ := strconv.Atoi(s)
		dataset[i] = s
	}
	return dataset
}
