package main

import (
	"fmt"
	"slices"
	// "strings"
)

func p1(data [][]int64, p1 bool) {
	sum := int64(0)

	for _, line := range data {
		if !p1 {
			slices.Reverse(line)
		}
		// txt := lineStr(line)
		// fmt.Printf("%s\n", txt)
		v := processLine(line, 1)
		sum += v
		// fmt.Println(v)
		// fmt.Println()
	}

	if p1 {
		fmt.Println("pt 1:", sum)
	} else {

		fmt.Println("pt 2:", sum)
	}
}

func processLine(line []int64, lvl int) int64 {
	if allZeroes(line) {
		return line[len(line)-1]
	}
	// diffTxt := strings.Repeat("  ", lvl)
	diffs := []int64{}
	for i := 1; i < len(line); i++ {
		d := line[i] - line[i-1]
		diffs = append(diffs, d)

		//logging
		// if line[i-1] > 9 && d < 10 {
		// 	diffTxt += " "
		// }
		// diffTxt += fmt.Sprintf("%d", d)
		// diffTxt += "  "
		//

	}
	// fmt.Println(diffTxt)

	return line[len(line)-1] + processLine(diffs, lvl+1)
}

func allZeroes(line []int64) bool {
	for _, n := range line {
		if n != 0 {
			return false
		}
	}
	return true
}
func lineStr(line []int64) string {
	txt := ""
	for i, elem := range line {
		txt += fmt.Sprintf("%d", elem)
		if i < len(line)-1 {
			txt += "  "
		}
	}
	return txt
}
