package day6

import (
	"errors"
	"fmt"
)

type DIR int

const (
	UP DIR = iota + 1
	DOWN
	LEFT
	RIGHT
)

func (d DIR) String() string {
	switch d {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	}
	return ""
}
func (d DIR) shape() rune {
	switch d {
	case UP:
		return '^'
	case DOWN:
		return 'v'
	case LEFT:
		return '<'
	case RIGHT:
		return '>'
	}
	return '.'
}

func Run(data []string) error {
	room := make([][]rune, len(data))
	startingDir := UP
	startPos := make([]int, 2)
	for i, l := range data {
		room[i] = []rune(l)
		for j, p := range room[i] {
			if p == '^' {
				startingDir = UP
				startPos[0] = i
				startPos[1] = j
			} else if p == 'v' {
				startingDir = DOWN
				startPos[0] = i
				startPos[1] = j
			} else if p == '<' {
				startingDir = LEFT
				startPos[0] = i
				startPos[1] = j
			} else if p == '>' {
				startingDir = RIGHT
				startPos[0] = i
				startPos[1] = j
			}
		}
	}
	for i, l := range room {
		fmt.Printf("%d: %c\n", i, l)
	}
	// startingDir = RIGHT
	err := p1(startPos, room, startingDir)
	if err != nil {
		return err
	}
	err = p2(startPos, room, startingDir)
	if err != nil {
		return err
	}
	return nil
}

func p1(startingPos []int, room [][]rune, direction DIR) error {
	visited := map[string]bool{}
	for i, j := startingPos[0], startingPos[1]; inbounds(i, j, room); i, j = update(i, j, direction) {
		if room[i][j] == '#' {
			i, j, direction = turn(i, j, direction)
		}
		visited[fmt.Sprint(i, j)] = true
	}

	ans := len(visited)
	fmt.Printf("Part 1: %d\n", ans)

	return nil
}

func p2(startingPos []int, room [][]rune, direction DIR) error {
	/*
		find the places s.t. placing an obstacle would cause the guard to travel along the a path she's already taken

		In the visited map, store the direction being traveled.
		If placing an object causes the guard to travel a path already taken, it'd be a loop
			that means, check if you can make the guard turn down that path
			If the gauard is going left, and there is a path traveling up next to them, you can place an obstacle in front of them


	*/
	return errors.New(`test all candidates independently
use p1`)

	startingDir := direction
	visited := map[string]DIR{}
	candidates := map[string]bool{}
	for i, j := startingPos[0], startingPos[1]; inbounds(i, j, room); i, j = update(i, j, direction) {
		if room[i][j] == '#' {
			i, j, direction = turn(i, j, direction)
		}
		visited[fmt.Sprint(i, j)] = direction
	}
	direction = startingDir
	for i, j := startingPos[0], startingPos[1]; inbounds(i, j, room); i, j = update(i, j, direction) {
		if room[i][j] == '#' {
			i, j, direction = turn(i, j, direction)
		}

		// check for candidate
		switch direction {
		case UP:
			//look if path is to the right going right
			if d, prs := visited[fmt.Sprint(i, j+1)]; prs && d == RIGHT {
				if i > 0 && room[i-1][j] != '#' {
					candidates[fmt.Sprint(i-1, j)] = true
				}
			}
		case DOWN:
			//look if path is to the left
			if d, prs := visited[fmt.Sprint(i, j-1)]; prs && d == LEFT {
				if j < len(room) && room[i+1][j] != '#' {
					candidates[fmt.Sprint(i+1, j)] = true
				}
			}
		case LEFT:
			//look if path is to up
			if d, prs := visited[fmt.Sprint(i-1, j)]; prs && d == UP {
				if j > 0 && room[i][j-1] != '#' {
					candidates[fmt.Sprint(i, j-1)] = true
				}
			}
		case RIGHT:
			//look if path is to down
			if d, prs := visited[fmt.Sprint(i+1, j)]; prs && d == DOWN {
				if j < len(room[0]) && room[i][j+1] != '#' {
					candidates[fmt.Sprint(i, j+1)] = true
				}
			}
		}
	}

	ans := len(candidates)
	fmt.Printf("Part 2: %d\n", ans)
	// return nil
	//done
	delete(visited, fmt.Sprint(startingPos[0], startingPos[1]))
	fmt.Println()
	for i := range room {
		row := []rune{}
		for j := range room[i] {
			k := fmt.Sprint(i, j)
			if _, prs := candidates[k]; prs {
				row = append(row, '')
			} else if d, prs := visited[k]; prs {
				// row = append(row, 'x')
				row = append(row, d.shape())
			} else {
				row = append(row, room[i][j])
			}
		}
		fmt.Printf("%d: %c", i, row)
		fmt.Println()
	}
	return nil
}

func inbounds(i, j int, room [][]rune) bool {
	withinX := i >= 0 && i < len(room)
	withinY := j >= 0 && j < len(room[0])
	return withinX && withinY
}
func update(i, j int, direction DIR) (int, int) {
	switch direction {
	case UP:
		return i - 1, j
	case DOWN:
		return i + 1, j
	case LEFT:
		return i, j - 1
	case RIGHT:
		return i, j + 1
	}
	return i, j
}
func turn(i, j int, dir DIR) (int, int, DIR) {
	switch dir {
	case UP:
		return i + 1, j, RIGHT
	case DOWN:
		return i - 1, j, LEFT
	case LEFT:
		return i, j + 1, UP
	case RIGHT:
		i, j = update(i, j, DOWN)
		return i, j - 1, DOWN
	}
	return i, j, UP
}
