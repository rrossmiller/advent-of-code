package day4

import (
	"fmt"
	"slices"
	"strings"
)

var (
	FORWARDS = []rune{'X', 'M', 'A', 'S'}
	// BACKWARDS = []rune{'S', 'A', 'M', 'X'}
	LEN  = len(FORWARDS)
	DIRS = [][]int{
		// up left up up right
		{-1, -1}, {-1, 0}, {-1, 1},
		//left , right
		{0, -1}, {0, 1},
		//down left, down, down right
		{1, -1}, {1, 0}, {1, 1},
	}

	ROWS = 0
	COLS = 0
	// antagonist matches don't double countt
	visitedUp   = []point{}
	visitedDown = []point{}

	visitedLeft  = []point{}
	visitedRight = []point{}

	visitedUpLeft    = []point{}
	visitedDownRight = []point{}

	visitedUpRight  = []point{}
	visitedDownLeft = []point{}
)

type point struct {
	x, y int
}

func Run(data []string) error {
	lines := make([][]rune, 0, len(data))
	for _, l := range data {
		lines = append(lines, []rune(l))
		// fmt.Println(l)
	}
	ROWS = len(lines)
	COLS = len(lines[0])

	p1(lines)
	return nil
}

func p1(lines [][]rune) error {
	ans := 0
	for i, l := range lines {
		for j, c := range l {
			if c == 'X' || c == 'S' {
				ans += search(i, j, lines)
			}
		}
	}
	fmt.Printf("Part 1: %d\n", ans)
	if ans >= 2477 {
		fmt.Println("too high")
	} else if ans <= 2179 {
		fmt.Println("too low")
	}

	if len(lines) < 100 {
		fmt.Println()
		for i, l := range lines {
			for j, c := range l {
				if isVisitedAny(point{i, j}) {
					fmt.Printf("%c", c)
				} else {
					s := strings.ToLower(string(c))
					fmt.Printf(s)
				}
			}
			fmt.Print("\n")
		}
		fmt.Println()

		x := strings.Split(`....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX`, "\n")
		for i, l := range lines {
			for j, c := range l {
				if isVisitedAny(point{i, j}) {
					// if isVisited(point{i, j}, visitedUp) {
					// if isVisited(point{i, j}, visitedDown) {
					// if isVisited(point{i, j}, visitedLeft) {
					// if isVisited(point{i, j}, visitedRight) {
					// if isVisited(point{i, j}, visitedUpLeft) {
					// if isVisited(point{i, j}, visitedDownRight) {
					// if isVisited(point{i, j}, visitedUpRight) {
					// if isVisited(point{i, j}, visitedDownLeft) {
					fmt.Printf("%c", c)
				} else {
					fmt.Print(".")
				}
			}
			fmt.Printf(" | %v", x[i])
			fmt.Print("\n")
		}
		fmt.Println()
		sort := func(a, b point) int {
			if a.x < b.x {
				return -1
			} else if a.x > b.x {
				return 1
			}

			return 0
		}
		slices.SortFunc(visitedUp, sort)
		slices.SortFunc(visitedDown, sort)
		slices.SortFunc(visitedLeft, sort)
		slices.SortFunc(visitedRight, sort)
		slices.SortFunc(visitedUpRight, sort)
		slices.SortFunc(visitedDownLeft, sort)
		slices.SortFunc(visitedUpLeft, sort)
		slices.SortFunc(visitedDownRight, sort)
		fmt.Printf("%v\n%v\n\n", visitedUp, visitedDown)
		fmt.Printf("%v\n%v\n\n", visitedLeft, visitedRight)
		fmt.Printf("%v\n%v\n\n", visitedUpLeft, visitedDownRight)
		fmt.Printf("%v\n%v\n\n", visitedUpRight, visitedDownLeft)
	}
	return nil
}

// func p2(data []string) error { return nil }

func search(i, j int, lines [][]rune) int {
	n := 0
	// check for room to search right and left
	if j >= LEN {
		n += searchLeft(i, j, lines)
	}
	if j <= COLS-LEN {
		n += searchRight(i, j, lines)
	}

	// if i is greater than the lenght of the word, there's space to search above
	if i >= LEN {
		n += searchUp(i, j, lines)
		// if j is greater than the lenght of the word, there's space to search up and left
		if j >= LEN-1 {
			n += searchUpLeft(i, j, lines)
		}
		// check for room to search up and right
		if j < COLS-LEN {
			n += searchUpRight(i, j, lines)
		}
	}

	// check for room to search down
	if i <= ROWS-LEN {
		n += searchDown(i, j, lines)
		if j > LEN {
			n += searchDownLeft(i, j, lines)
		}
		if j < COLS-LEN {
			n += searchDownRight(i, j, lines)
		}
	}
	return n
}

func isVisitedAny(pnt point) bool {
	visited := append(visitedUp, visitedDown...)
	visited = append(visited, visitedLeft...)
	visited = append(visited, visitedRight...)

	visited = append(visited, visitedUpLeft...)
	visited = append(visited, visitedDownRight...)

	visited = append(visited, visitedUpRight...)
	visited = append(visited, visitedDownLeft...)
	for _, p := range visited {
		if p.x == pnt.x && p.y == pnt.y {
			return true
		}
	}

	return false
}
func isVisited(pnt point, visited []point) bool {
	for _, p := range visited {
		if p.x == pnt.x && p.y == pnt.y {
			return true
		}
	}

	return false
}
func searchUp(i, j int, lines [][]rune) int {
	var sb strings.Builder
	points := make([]point, 0, LEN)
	for x := range LEN {
		sb.WriteRune(lines[i-x][j])
		points = append(points, point{i - x, j})
	}

	if isVisited(points[len(points)-2], visitedDown) {
		return 0
	}
	if sb.String() == string(FORWARDS) { //|| sb.String() == string(BACKWARDS) {
		visitedUp = append(visitedUp, points...)
		return 1
	}
	return 0
}

func searchDown(i, j int, lines [][]rune) int {
	var sb strings.Builder
	points := make([]point, 0, LEN)
	for x := range LEN {
		sb.WriteRune(lines[i+x][j])
		points = append(points, point{i + x, j})
	}

	if isVisited(points[len(points)-2], visitedUp) {
		return 0
	}
	if sb.String() == string(FORWARDS) { //|| sb.String() == string(BACKWARDS) {
		visitedDown = append(visitedDown, points...)
		return 1
	}
	return 0
}
func searchLeft(i, j int, lines [][]rune) int {
	var sb strings.Builder
	points := make([]point, 0, LEN)
	for x := range LEN {
		sb.WriteRune(lines[i][j-x])
		points = append(points, point{i, j - x})
	}

	if isVisited(points[len(points)-2], visitedRight) {
		return 0
	}
	if sb.String() == string(FORWARDS) { //|| sb.String() == string(BACKWARDS) {
		visitedLeft = append(visitedLeft, points...)
		return 1
	}
	return 0
}
func searchRight(i, j int, lines [][]rune) int {
	var sb strings.Builder
	points := make([]point, 0, LEN)
	for x := range LEN {
		sb.WriteRune(lines[i][j+x])
		points = append(points, point{i, j + x})
	}

	if isVisited(points[len(points)-2], visitedLeft) {
		return 0
	}
	if sb.String() == string(FORWARDS) { //|| sb.String() == string(BACKWARDS) {
		visitedRight = append(visitedRight, points...)
		return 1
	}
	return 0
}
func searchUpLeft(i, j int, lines [][]rune) int {
	var sb strings.Builder
	points := make([]point, 0, LEN)
	for x := range LEN {
		sb.WriteRune(lines[i-x][j-x])
		points = append(points, point{i - x, j - x})
	}

	if isVisited(points[len(points)-2], visitedDownRight) {
		return 0
	}
	if sb.String() == string(FORWARDS) { //|| sb.String() == string(BACKWARDS) {
		visitedUpLeft = append(visitedUpLeft, points...)
		return 1
	}
	return 0
}
func searchUpRight(i, j int, lines [][]rune) int {
	var sb strings.Builder
	points := make([]point, 0, LEN)
	for x := range LEN {
		sb.WriteRune(lines[i-x][j+x])
		points = append(points, point{i - x, j + x})
	}

	if isVisited(points[len(points)-2], visitedDownLeft) {
		return 0
	}
	if sb.String() == string(FORWARDS) { //|| sb.String() == string(BACKWARDS) {
		visitedUpRight = append(visitedUpRight, points...)
		return 1
	}
	return 0
}
func searchDownLeft(i, j int, lines [][]rune) int {
	var sb strings.Builder
	points := make([]point, 0, LEN)
	for x := range LEN {
		sb.WriteRune(lines[i+x][j-x])
		points = append(points, point{i + x, j - x})
	}

	if isVisited(points[len(points)-2], visitedUpRight) {
		return 0
	}
	if sb.String() == string(FORWARDS) { //|| sb.String() == string(BACKWARDS) {
		visitedDownLeft = append(visitedDownLeft, points...)
		return 1
	}
	return 0
}
func searchDownRight(i, j int, lines [][]rune) int {
	var sb strings.Builder
	points := make([]point, 0, LEN)
	for x := range LEN {
		sb.WriteRune(lines[i][j+x])
		points = append(points, point{i + x, j + x})
	}

	if isVisited(points[len(points)-2], visitedUpLeft) {
		return 0
	}
	if sb.String() == string(FORWARDS) { //|| sb.String() == string(BACKWARDS) {
		visitedDownRight = append(visitedDownRight, points...)
		return 1
	}
	return 0
}
