package main

import (
	"aoc/datas"
	"aoc/day1"
	"aoc/day2"
	"flag"
	"fmt"
	"path/filepath"
	"regexp"
	"slices"
)

var dayRegex = regexp.MustCompile(`.*(\d+)`)

func main() {
	var test bool
	var day string
	flag.BoolVar(&test, "t", false, "Run with test data")
	flag.StringVar(&day, "d", "", "Run a specific day")

	flag.Parse()

	if day == "" {
		today, err := getLatestDay()
		if err != nil {
			panic(err)
		}
		day = today
	}
	fmt.Printf("Running day %s\n\n", day)

	data := datas.GetData(day+".txt", test)

	var err error
	switch day {
	case "1":
		err = day1.Run(data)
	case "2":
		err = day2.Run(data)
	}
	if err != nil {
		panic(err)
	}
}

// get the number of the current day (based on what days are in the data dir)
func getLatestDay() (string, error) {
	matches, err := filepath.Glob("../data/*")
	if err != nil {
		return "", nil
	}

	slices.Sort(matches)
	day := matches[len(matches)-1]

	days := dayRegex.FindAllStringSubmatch(day, -1)
	return days[0][1], nil
}