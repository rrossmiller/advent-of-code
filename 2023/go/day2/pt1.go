package main

import (
	"fmt"
	"strconv"
	"strings"
)

func pt1(data string) {
	validGameSum := 0
	// parse rows into a list of games
	for i, r := range strings.Split(data, "\n") {
		if len(r) == 0 {
			continue
		}

		game := strings.Split(r, ":")
		validGame := true
		for _, rnd := range strings.Split(game[1], ";") {
			for _, pick := range strings.Split(rnd, ",") {
				x := strings.Split(strings.TrimSpace(pick), " ")
				color := x[1]
				cnt, err := strconv.ParseInt(x[0], 10, 32)
				if err != nil {
					panic(err)
				}
				// check that the colors are valid
				// check that the num picked is less than the num available
				if v, exist := colors[color]; !exist || cnt > int64(v) {
					validGame = false
					break
				}
			}
		}

		if validGame {
			validGameSum += i + 1
		}

	}

	fmt.Println("pt1:", validGameSum)
}
