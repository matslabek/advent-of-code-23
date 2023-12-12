package wait

import (
	trebuchet "AdventOfCode/1"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const FilePath = "6/input.txt"

func WaitForIt() {
	inputStrings, err := trebuchet.ReadStringsFromFile(FilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	times := strings.Fields(inputStrings[0])[1:]
	distances := strings.Fields(inputStrings[1])[1:]
	// Part 1
	product := 1
	for i := 0; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		ww := calculateWinningWays(t, d)
		product *= 1 + ww[1] - ww[0]
	}
	fmt.Println(product)

	//Part 2
	time := strings.Join(times[:], "")
	dist := strings.Join(distances[:], "")
	t2, _ := strconv.Atoi(time)
	d2, _ := strconv.Atoi(dist)

	ww2 := calculateWinningWays(t2, d2)
	fmt.Println(ww2)
	prod2 := 1 + ww2[1] - ww2[0]
	fmt.Println(prod2)
}

func calculateWinningWays(time, distanceToBeat int) []int {
	// All we need to do is calculate zeros of a quadratic function
	x1 := (float64(time) - math.Sqrt(float64(time*time-(4*distanceToBeat)))) / 2
	x2 := (float64(time) + math.Sqrt(float64(time*time-(4*distanceToBeat)))) / 2
	// Edge cases
	if x1 == math.Ceil((x1)) {
		x1++
	}
	if x2 == math.Floor(x2) {
		x2--
	}
	return []int{int(math.Ceil(x1)), int(math.Floor(x2))}
}
