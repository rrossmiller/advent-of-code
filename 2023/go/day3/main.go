package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	P1 = iota
	P2
)

func main() {
	data := getData()
	if len(os.Args) > 1 {
		data = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	}
	run(data, P1)
	run(data, P2)

}

func run(data string, part uint) {
	lines := strings.Split(data, "\n")
	nums := []int{}
	visited := []Coord{}

	dat := [][]rune{}
	for _, l := range lines {
		dat = append(dat, []rune(l))
		// fmt.Println(l)
	}

	// for every line
	for i, l := range lines {
		//for every rune in the line
		for j, r := range []rune(l) {
			// if the char is a symbol, look around it for numbers
			if !unicode.IsDigit(r) && r != '.' {
				switch part {
				case P1:
					nums, visited = p1(dat, i, j, nums, visited)
				case P2:
					if r == '*' {
						nums, visited = p2(dat, i, j, nums, visited)
					}
				}
			}
		}
	}

	s := 0
	// fmt.Println(nums)
	// fmt.Println(visited)
	for _, n := range nums {
		s += n
	}
	switch part {
	case P1:
		fmt.Printf("p1: %d\n", s)
	case P2:
		fmt.Printf("p2: %d\n", s)
	}

}

func getData() string {
	b, err := os.ReadFile("../../data/3.txt")
	if err != nil {
		panic(err)
	}
	data := string(b)

	return data
}
