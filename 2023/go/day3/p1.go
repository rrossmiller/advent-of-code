package main

import (
	"strconv"
	"unicode"
)

type Coord struct {
	x    int
	y    int
	nums *[]int
}

const SYMBOLS = "!@#$%^&*()+-="

var (
	dirs = [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
)

func p1(dat [][]rune, i, j int, nums []int, visited []Coord) ([]int, []Coord) {
	cols := len(dat[0])
	// bfs to look around the symbol
	for _, d := range dirs {
		row := i + d[0]
		col := j + d[1]
		if row < 0 || row >= len(dat) || col < 0 || col >= cols {
			continue
		}
		coord := Coord{row, col, nil}
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
			nums = append(nums, n)
		}
	}
	return nums, visited
}

func isVisited(c Coord, s []Coord) bool {
	for _, p := range s {
		if p.x == c.x && p.y == c.y {
			return true
		}
	}
	return false
}
