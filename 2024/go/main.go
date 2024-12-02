package main

import (
	"aoc/day1"
	"flag"
)

func main() {
	var test bool
	flag.BoolVar(&test, "t", false, "")
	flag.Parse()
	day1.Run(test)
}

func getLatestDay() {

}
