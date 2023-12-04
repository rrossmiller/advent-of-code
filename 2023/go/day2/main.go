package main

import (
	"fmt"
	"os"
)

var (
	colors = map[string]int{"red": 12, "green": 13, "blue": 14}
)

type Game struct {
	gameNum int
}

func main() {
	fmt.Println("Day 2")
	data := getData()
	pt1(data)
	pt2(data)
}

func getData() string {
	b, err := os.ReadFile("../data/2.txt")
	if err != nil {
		panic(err)
	}
	data := string(b)

	return data

}
