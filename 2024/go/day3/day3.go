package day3

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	mulRegex  = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doRegex   = regexp.MustCompile(`do\(\)`)
	dontRegex = regexp.MustCompile(`don't\(\)`)
)

func Run(data []string) error {
	err := p1(data)
	if err != nil {
		return err
	}

	err = p2(data)
	if err != nil {
		return err
	}
	return nil
}

func p1(data []string) error {
	sum := 0
	for _, line := range data {
		// find instances of mul(n,n) in the line
		matches := mulRegex.FindAllStringSubmatch(line, -1)
		for _, m := range matches {
			ans, err := product(m[1:3])
			if err != nil {
				return err
			}
			sum += ans

		}
	}
	fmt.Printf("Part 1: %d\n", sum)
	return nil
}

func p2(data []string) error {
	sum := 0
	op := true
	for _, line := range data {
		// find instances of mul, do and don't in the line
		mul := mulRegex.FindAllStringSubmatchIndex(line, -1)
		do := doRegex.FindAllStringIndex(line, -1)
		dont := dontRegex.FindAllStringIndex(line, -1)
		matches := make([][]int, 0, len(mul)+len(do)+len(dont))
		matches = append(matches, mul...)
		matches = append(matches, do...)
		matches = append(matches, dont...)
		// put them in order
		slices.SortFunc(matches, func(a, b []int) int {
			if a[0] < b[0] {
				return -1
			} else if a[0] > b[0] {
				return 1
			}
			return 0
		})
		for _, m := range matches {
			phrase := line[m[0]:m[1]]
			// only get the product if you're in a do section
			if op && strings.Contains(phrase, "mul") {
				a := line[m[2]:m[3]]
				b := line[m[4]:m[5]]
				ans, err := product([]string{a, b})
				if err != nil {
					return err
				}
				sum += ans
			} else if phrase == "do()" {
				op = true
			} else if phrase == "don't()" {
				op = false
			}
		}
	}
	fmt.Printf("Part 2: %d\n", sum)
	return nil
}

func product(m []string) (int, error) {
	a, err := strconv.Atoi(m[0])
	if err != nil {
		return 0, err
	}
	b, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, err
	}
	return a * b, nil
}
