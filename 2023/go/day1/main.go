package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	data := readFile(false)
	start := time.Now()
	p1(data)
	end := time.Since(start)
	fmt.Printf("Elapsed (chan): %v\n", end)

	start = time.Now()
	p1a(data)
	end = time.Since(start)
	fmt.Printf("Elapsed (atomic): %v\n", end)

	fmt.Println()
// 	data = `two1nine
// eightwothree
// abcone2threexyz
// xtwone3four
// 4nineeightseven2
// zoneight234
// 7pqrstsixteen`
	p2(data)
}

func readFile(test bool) string {
	if test {
		return `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	}
	b, err := os.ReadFile("../../data/1.txt")
	if err != nil {
		panic(err)
	}

	return string(b)
}
