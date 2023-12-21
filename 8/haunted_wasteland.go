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
	//instructions := inputStrings[0]
	n := len(inputStrings)
	m := make(map[string][]string)

	for i := 2; i < n; i++ {
		s := strings.Fields(inputStrings[i])
		mm := make([]string, 2)
		mm = append(mm, s[2][1:4], s[3][0:3])
		m[s[0]] = mm
	}
	fmt.Println(m)
}
