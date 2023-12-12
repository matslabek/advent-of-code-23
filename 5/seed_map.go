package seed

import (
	trebuchet "AdventOfCode/1"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const FilePath = "5/input.txt"

func Seed() {

	almanacSlice := make([][][]int, 0)

	inputStrings, err := trebuchet.ReadStringsFromFile(FilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	s := make([][]int, 0)
	// Prepare the slices
	for i := 2; i < len(inputStrings); i++ {
		if inputStrings[i] == "" {
			// This is because of the unexpected bug that arose from 's' capacity being exceeded\
			// Previous version of the code:
			//		almanacSlice = append(almanacSlice, s)
			// caused almanacSlice to have its values overwritten when s got overflowed.
			almanacSlice = append(almanacSlice, make([][]int, len(s)))
			copy(almanacSlice[len(almanacSlice)-1], s)
			// Clear the original slice
			s = s[:0]
		} else {
			firstChar := inputStrings[i][0]
			_, err := strconv.Atoi(string(firstChar))
			if err == nil {
				splitRes := strings.Fields(inputStrings[i])
				source, _ := strconv.Atoi(splitRes[1])
				destination, _ := strconv.Atoi(splitRes[0])
				rangeLen, _ := strconv.Atoi(splitRes[2])
				start := source
				end := source + rangeLen
				shift := destination - source
				ss := make([]int, 0)
				ss = append(ss, start, end, shift)
				s = append(s, ss)
			}
		}
	}
	noOfMaps := 7
	// Find the value
	finalValues := make([]int, 0)
	seeds := strings.Fields(strings.Split(inputStrings[0], ":")[1])
	sI, _ := stringsToIntegers(seeds)
	seedRanges := prepareSeedRange(sI)
	fmt.Println(seedRanges)
	// Part 1
	for i := 0; i < len(seeds); i++ {
		seedValue, _ := strconv.Atoi(seeds[i])
		for j := 0; j < noOfMaps; j++ {
			for src, destRg := range almanacSlice[j] {
				dest := destRg[0]
				rg := destRg[1]
				if seedValue >= src && seedValue <= src+rg {
					seedValue = dest + (seedValue - src)
					break
				}
			}
		}
		finalValues = append(finalValues, seedValue)
	}
	fmt.Println(finalValues)
	// Final answer part 1
	fmt.Println(slices.Min(finalValues))

	//Part two:
	sortedMaps := sortAlmanacSlice(almanacSlice)
	fmt.Println(sortedMaps)
	//TODO: make this work
	for i := 0; i < len(seedRanges); i++ {
		fmt.Println("All seed ranges:", seedRanges)
		for j := 0; j < noOfMaps; j++ {
			for k := 0; k < len(sortedMaps[j]); k++ {
				currentRange := seedRanges[i]
				fmt.Println("Current seed range", currentRange)
				// Start falls within range
				fmt.Println("Current map range", sortedMaps[j][k])
				if seedRanges[i][0] >= sortedMaps[j][k][0] && seedRanges[i][0] <= sortedMaps[j][k][1] {
					seedRanges[i][0] += sortedMaps[j][k][2]
					// End falls in the range too
					if seedRanges[i][1] >= sortedMaps[j][k][0] && seedRanges[i][1] <= sortedMaps[j][k][1] {
						seedRanges[i][1] += sortedMaps[j][k][2]
						break
					} else {
						// End doesn't fall in the range - create new slice
						s := make([]int, 2)
						// End of mapRange is new start, end of seedRange is new end
						s[0] = sortedMaps[j][k][1] + 1
						s[1] = seedRanges[i][1]
						seedRanges[i][1] = sortedMaps[j][k][1] + sortedMaps[j][k][2]
						seedRanges = append(seedRanges, s)
						break
					}
				} else if seedRanges[i][0] < sortedMaps[j][k][0] {
					// Start doesn't fall within the range
					// End falls:
					if seedRanges[i][1] >= sortedMaps[j][k][0] && seedRanges[i][1] <= sortedMaps[j][k][1] {
						s := make([]int, 2)
						// start of mapRange is new start, end of seedRange is new end
						s[0] = sortedMaps[j][k][0]
						s[1] = seedRanges[i][1]
						seedRanges = append(seedRanges, s)
						break
					} else {
						// End is lower than the range start, break
						if seedRanges[i][1] < sortedMaps[j][k][0] {
							break
						}
						// End is higher than the range end, then cut it into two slices.
						s := make([]int, 2)
						s[0] = sortedMaps[j][k][0]
						s[1] = sortedMaps[j][k][1]

						seedRanges = append(seedRanges, s)

						r := make([]int, 2)
						r[0] = sortedMaps[j][k][1] + 1
						r[1] = seedRanges[i][1]

						seedRanges = append(seedRanges, r)
						seedRanges[i][1] = sortedMaps[j][k][0] - 1
						break
					}
				}
			}
		}
	}
	fmt.Println(seedRanges)

}
func prepareSeedRange(seeds []int) [][]int {
	sr := make([][]int, 0)
	for i := 1; i < len(seeds); i += 2 {
		s := make([]int, 0)
		s = append(s, seeds[i-1], seeds[i-1]+seeds[i])
		sr = append(sr, s)
	}
	return sr
}

func stringsToIntegers(lines []string) ([]int, error) {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}

func sortAlmanacSlice(s [][][]int) [][][]int {

	// Define a custom comparator function for sorting
	comparator := func(i, j int) bool {
		return s[i][0][0] < s[j][0][0]
	}

	// Use sort.Slice to sort the s based on the comparator function
	sort.Slice(s, comparator)

	for _, innerSlice := range s {
		sort.Slice(innerSlice, func(i, j int) bool {
			return innerSlice[i][0] < innerSlice[j][0]
		})
	}
	return s
}
