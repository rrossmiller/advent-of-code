package main

import (
	"fmt"
	"strings"
)

func p1(dat []string) {
	lines := [][]rune{}
	for _, s := range dat {
		inner := []rune{}
		for _, r := range s {
			inner = append(inner, r)
		}
		lines = append(lines, inner)
	}

	// // DELETEME
	// for _, l := range lines {
	// 	fmt.Printf("%c\n", l)
	// }

	// loop through rows
	for i := 1; i < len(lines); i++ {
		for j, c := range lines[i] {
			// if row above is a # or 0, move it up --> update in place
			if c == 'O' && lines[i-1][j] == '.' {
				lines[i][j] = '.'

				// move up until a stop char is encountered
				row := i
				for row > 0 && lines[row-1][j] == '.' {
					row -= 1
				}
				lines[row][j] = 'O'
			}
		}
	}

	// DELETEME
	// fmt.Println()
	// for _, l := range lines {
	// 	fmt.Printf("%c\n", l)
	// }
	// check(lines)
	// return

	// sum rows
	sum := 0
	for i, row := range lines {
		numO := 0
		for _, c := range row {
			if c == 'O' {
				numO += 1
			}
		}
		sum += numO * (len(lines) - i)
	}
	fmt.Printf("p1: %d\n", sum)
	fmt.Println("too high", 107918)
}

func check(lines [][]rune) {
	a := `OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`
	b := strings.Split(a, "\n")

	t := [][]rune{}
	for _, s := range b {
		inner := []rune{}
		for _, r := range s {
			inner = append(inner, r)
		}
		t = append(t, inner)
	}

	// DELETEME
	fmt.Println("asdf")
	for _, l := range t {
		fmt.Printf("%c\n", l)
	}
	for i, l := range lines {
		for j, r := range l {
			if x := t[i][j]; x != r {
				fmt.Printf("%d %d\n", i, j)
				panic("")
			}
		}
	}
}
