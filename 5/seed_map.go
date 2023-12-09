package seed

import (
	trebuchet "AdventOfCode/1"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const FilePath = "5/input.txt"

func Seed() {

	almanacMap := make(map[string]map[int][]int)

	inputStrings, err := trebuchet.ReadStringsFromFile(FilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	mapType := ""
	// Prepare the maps
	for i := 1; i < len(inputStrings); i++ {
		if inputStrings[i] == "" {
			continue
		}
		splitRes := strings.Split(inputStrings[i], "-")
		if len(splitRes) > 1 {
			mapType = splitRes[0] + splitRes[1] + strings.Fields(splitRes[2])[0]

			// Initialize the inner map if it doesn't exist
			if almanacMap[mapType] == nil {
				almanacMap[mapType] = make(map[int][]int)
			}
		} else if len(splitRes[0]) > 0 {
			splitRes = strings.Fields(splitRes[0])
			source, _ := strconv.Atoi(splitRes[1])
			destination, _ := strconv.Atoi(splitRes[0])
			rangeLen, _ := strconv.Atoi(splitRes[2])
			ss := make([]int, 0)
			ss = append(ss, destination)
			ss = append(ss, rangeLen)
			almanacMap[mapType][source] = ss

		}
	}
	mapNames := [7]string{"seedtosoil", "soiltofertilizer",
		"fertilizertowater", "watertolight", "lighttotemperature",
		"temperaturetohumidity", "humiditytolocation",
	}

	// Find the value
	finalValues := make([]int, 0)
	seeds := strings.Fields(strings.Split(inputStrings[0], ":")[1])

	for i := 0; i < len(seeds); i++ {
		seedValue, _ := strconv.Atoi(seeds[i])
		for j := 0; j < len(mapNames); j++ {
			for src, destRg := range almanacMap[mapNames[j]] {
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
	// Final answer
	fmt.Print(slices.Min(finalValues))
}
