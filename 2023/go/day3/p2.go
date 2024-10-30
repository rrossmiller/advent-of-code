package main

import (
	"strconv"
	"unicode"
)

func p2(dat [][]rune, i, j int, nums []int, visited []Coord) ([]int, []Coord) {
	cols := len(dat[0])
	foundNums := []int{}
	// bfs to look around the symbol
	for _, d := range dirs {
		row := i + d[0]
		col := j + d[1]
		if row < 0 || row >= len(dat) || col < 0 || col >= cols {
			continue
		}
		coord := Coord{row, col, &[]int{}}
		line := dat[row]
		// if the char at the coord is a digit
		// idx hasn't been checked already
		if unicode.IsDigit(line[col]) && !isVisited(coord, visited) {
			// visit the coord
			visited = append(visited, coord)
			// find the rest of the number
			l := col - 1
			r := col + 1
			// find leftmost limit
			for l > 0 && unicode.IsDigit(line[l]) {
				visited = append(visited, Coord{row, l, nil})
				l -= 1
			}
			if l < 0 || !unicode.IsDigit(line[l]) {
				l += 1
			}
			// find rightmost limit
			for r < len(line) && unicode.IsDigit(line[r]) {
				visited = append(visited, Coord{row, r, nil})
				r += 1
			}

			n, err := strconv.Atoi(string(line[l:r]))
			if err != nil {
				panic(err)
			}
			foundNums = append(foundNums, n)
		}
	}

	if len(foundNums) == 2 {
		nums = append(nums, foundNums[0]*foundNums[1])
	}
	return nums, visited
}
