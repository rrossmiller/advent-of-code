package datas

import (
	"os"
	"strings"
)

func GetData(day string, test bool) []string {
	b, _ := os.ReadFile("../data/" + day)
	lines := strings.Split(string(b), "\n")
	testLine := 0
	for i := range len(lines) {
		if lines[i] == "-----TEST-----" {
			testLine = i
			break
		}
	}

	if test {
		return lines[:testLine]
	}

	return lines[testLine+1:]
}
