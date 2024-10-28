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

func p1a(data string) {
	// create a channel with as many spaces as lines
	lines := strings.Split(data, "\n")
	var sum atomic.Int64
	var wg sync.WaitGroup

	// fill the channel
	for _, line := range lines {
		if utf8.RuneCountInString(line) > 0 {
			wg.Add(1)
			go aprocessLine(line, &sum, &wg)
		}
	}
	wg.Wait()
	fmt.Printf("pt1: %d\n", sum.Load())
}

func p1(data string) {
	// create a channel with as many spaces as lines
	lines := strings.Split(data, "\n")
	ch := make(chan int, len(lines))
	var wg sync.WaitGroup

	// fill the channel
	for _, line := range lines {
		if utf8.RuneCountInString(line) > 0 {
			wg.Add(1)
			go processLine(line, ch, &wg)
		}
	}
	wg.Wait()
	close(ch)
	// sum accross the channel
	sum := 0
	for i := range ch {
		sum += i
	}
	fmt.Printf("pt1: %d\n", sum)
}

func aprocessLine(line string, sum *atomic.Int64, wg *sync.WaitGroup) {
	defer wg.Done()
	var sb strings.Builder

	// find first digit
	for _, c := range line {
		if unicode.IsDigit(c) {
			sb.WriteRune(c)
			break
		}
	}
	// find last digit
	for i := utf8.RuneCountInString(line) - 1; i >= 0; i-- {
		c := rune(line[i])
		if unicode.IsDigit(c) {
			sb.WriteRune(c)
			break
		}
	}

	n, err := strconv.ParseInt(sb.String(), 10, 64)
	if err != nil {
		panic(err)
	}
	sum.Add(n)
}
func processLine(line string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var sb strings.Builder

	// find first digit
	for _, c := range line {
		if unicode.IsDigit(c) {
			sb.WriteRune(c)
			break
		}
	}
	// find last digit
	for i := utf8.RuneCountInString(line) - 1; i >= 0; i-- {
		c := rune(line[i])
		if unicode.IsDigit(c) {
			sb.WriteRune(c)
			break
		}
	}

	n, err := strconv.ParseInt(sb.String(), 10, 64)
	if err != nil {
		panic(err)
	}
	ch <- int(n)
}
