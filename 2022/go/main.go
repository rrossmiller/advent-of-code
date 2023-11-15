package main

import (
	"aoc/days"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must pass in the number of the day to run (i.e., '1')")
		os.Exit(1)
	}

	if os.Args[1] == "1" {
		days.One()
	} else if os.Args[1] == "7" {
		days.Seven(100_000)
	} else if os.Args[1] == "8" {
		days.Eight()
	} else if os.Args[1] == "11" {
		days.Day11()
	} else if os.Args[1] == "12" {
		days.Day12()
	}
}
