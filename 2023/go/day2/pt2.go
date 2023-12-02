package main

import (
	"fmt"
	"strconv"
	"strings"
)

func pt2(data string) {
	validGameSum := 0
	// parse rows into a list of games
	for _, r := range strings.Split(data, "\n") {
		if len(r) == 0 {
			continue
		}
		game := strings.Split(r, ":")
		cubes := make(map[string]int64)
		for _, rnd := range strings.Split(game[1], ";") {
			for _, pick := range strings.Split(rnd, ",") {
				x := strings.Split(strings.TrimSpace(pick), " ")
				color := x[1]
				cnt, err := strconv.ParseInt(x[0], 10, 32)
				if err != nil {
					panic(err)
				}
				// if the color hasn't been recorded, or the count is greater than what's stored
				if pv, pExist := cubes[color]; !pExist || cnt > pv {
					cubes[color] = cnt
				}

			}
		}

		// sum powers
		power := 1
		for _, v := range cubes {
			power *= int(v)
		}

		validGameSum += power

	}

	fmt.Println("pt2:", validGameSum)
}
