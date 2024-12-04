package datas

import (
	"os"
	"strings"
)

func GetData(day string, test bool) ([]string, error) {
	b, err := os.ReadFile("../data/" + day)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")
	testLine := 0
	for i := range len(lines) {
		if lines[i] == "-----TEST-----" {
			testLine = i
			break
		}
	}

	if test {
		return lines[:testLine], nil
	}

	return lines[testLine+1 : len(lines)-1], nil
}
