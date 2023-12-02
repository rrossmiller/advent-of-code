package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	colors = map[string]int{"red": 12, "green": 13, "blue": 14}
)

type Game struct {
	gameNum int
}

func main() {
	data := getData()
	pt1(data)

}

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
		// break

	}

	fmt.Println("pt1:", validGameSum)
}
func getData() string {
	b, err := os.ReadFile("../data/2.txt")
	if err != nil {
		panic(err)
	}
	data := string(b)

	return data

}
