package cube_conundrum

import (
	"AdventOfCode/1"
	"fmt"
	"strconv"
	"strings"
)

func Cube() {
	const FilePath = "2/input.txt"
	const RedMax = 12
	const GreenMax = 13
	const BlueMax = 14

	strings, err := trebuchet.ReadStringsFromFile(FilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	sum := 0
	// Count the sum
	for _, str := range strings {
		gameResult := parseCubeString(str)
		for gameNumber, result := range gameResult {
			if result["red"] <= RedMax && result["green"] <= GreenMax && result["blue"] <= BlueMax {
				sum += gameNumber
			}
		}
	}
	fmt.Println(sum)
}

func parseCubeString(gameResultString string) map[int]map[string]int {
	gameMap := make(map[string]int)
	gameMap["red"] = 0
	gameMap["blue"] = 0
	gameMap["green"] = 0

	splitResults := strings.Split(gameResultString, ":")
	gameNrString := splitResults[0]
	resultsString := splitResults[1]
	gameNumber, _ := strconv.Atoi(strings.Split(gameNrString, " ")[1])
	gameRounds := strings.Split(resultsString, ";")
	for _, round := range gameRounds {
		round = strings.TrimSpace(round)
		colors := strings.Split(round, ",")
		for _, color := range colors {
			color = strings.TrimSpace(color)
			results := strings.Split(color, " ")
			numberOfCubes, _ := strconv.Atoi(results[0])
			colorName := results[1]
			if gameMap[colorName] < numberOfCubes {
				gameMap[colorName] = numberOfCubes
			}
		}
	}
	gameResultMap := make(map[int]map[string]int)
	gameResultMap[gameNumber] = gameMap
	return gameResultMap
}
