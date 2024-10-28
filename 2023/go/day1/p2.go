package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"unicode"
	"unicode/utf8"
)

var (
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	keys = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func p2(data string) {
	// create a channel with as many spaces as lines
	lines := strings.Split(data, "\n")
	var sum atomic.Int64
	var wg sync.WaitGroup

	// fill the channel
	for _, line := range lines {
		if utf8.RuneCountInString(line) > 0 {
			wg.Add(1)
			// go processLine2(line, &sum, &wg)
			processLine2(line, &sum, &wg)
		}
	}
	wg.Wait()
	fmt.Printf("pt2: %d\n", sum.Load())
}

func processLine2(line string, sum *atomic.Int64, wg *sync.WaitGroup) {
	defer wg.Done()
	var sb strings.Builder

	// find first digit
	for i, c := range line {
		if unicode.IsDigit(c) {
			sb.WriteRune(c)
			break
		} else if n, ok := hasNum(line[0 : i+1]); ok {
			sb.WriteString(fmt.Sprint(n))
			break
		}
	}
	// find last digit
	for i := utf8.RuneCountInString(line) - 1; i >= 0; i-- {
		c := rune(line[i])
		if unicode.IsDigit(c) {
			sb.WriteRune(c)
			break
		} else if n, ok := hasNum(line[i:]); ok {
			sb.WriteString(fmt.Sprint(n))
			break
		}
	}

	n, err := strconv.ParseInt(sb.String(), 10, 64)
	if err != nil {
		panic(err)
	}
	sum.Add(n)
}

func hasNum(s string) (int, bool) {
	for i, k := range keys {
		if strings.Contains(s, k) {
			return nums[i], true
		}
	}
	return 0, false
}
