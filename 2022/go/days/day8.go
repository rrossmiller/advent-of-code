package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Eight() {
	/*
	   Limitation:
	   must be concurrent.
	   must not duplicate `dat`. Each thread must read from the same instance of dat

	   Challange:
	   each thread should write to the same result object (use mutex?)

	*/

	// how many trees are visible from outside the grid?
	dat := read8()

	rows := [][]int{}
	for _, r := range strings.Split(dat, "\n") {
		if len(r) == 0 {
			continue
		}
		inner := []int{}
		for _, i := range strings.Split(r, "") {
			// r = strings.TrimSpace(r)
			x, err := strconv.Atoi(i)
			if err != nil {
				// panic(err)
				continue
			}
			inner = append(inner, x)
		}
		rows = append(rows, inner)
	}
	//* * *
	//*   *
	//* * *
	// top+bottom + (left+right) -  4corners (don't double count the corners)
	visibleTrees := len(rows[0])*2 + len(rows)*2 - 4
	maxScore := 0

	for i := 1; i < len(rows)-1; i++ {
		for j := 1; j < len(rows[i])-1; j++ {
			elem := rows[i][j]

			canSeeLeft := true
			canSeeRight := true
			canSeeUp := true
			canSeeDown := true

			// scores
			upScore := 0
			downScore := 0
			leftScore := 0
			rightScore := 0

			// check left
			// can the tree see to the boarder
			for k := j - 1; k >= 0; k-- {
				leftScore++
				if elem <= rows[i][k] {
					canSeeLeft = false
					break
				}
			}

			// check right
			// can the tree see to the boarder
			for k := j + 1; k < len(rows[i]); k++ {
				rightScore++
				if elem <= rows[i][k] {
					canSeeRight = false
					break
				}
			}

			// check above
			// can the tree see to the boarder
			for k := i - 1; k >= 0; k-- {
				upScore++
				if elem <= rows[k][j] {
					canSeeUp = false
					break
				}
			}

			// check below
			// can the tree see to the boarder
			for k := i + 1; k < len(rows); k++ {
				downScore++
				if elem <= rows[k][j] {
					canSeeDown = false
					break
				}
			}

			if canSeeLeft || canSeeRight || canSeeUp || canSeeDown {
				visibleTrees++
			}

			score := upScore * downScore * leftScore * rightScore
			if score > maxScore {
				maxScore = score
			}
		}

	}
	fmt.Println(visibleTrees)
	fmt.Println(maxScore)

}

func read8() string {
	if false {
		dat := `30373
25512
65332
33549
35390`
		return dat

	}
	fBytes, err := os.ReadFile("../data/8.txt")
	Check(err)

	return string(fBytes)
}
