package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	// "time"
)

type Card struct {
	id             int
	winningNumbers []int
	cardNumbers    []int
}

func main() {
	fmt.Println("*****")
	fmt.Println("Day 4")
	data := getData()
	// data := getTestData()
	p1 := pt1(data)
	fmt.Println("Pt1:", p1)

	// parallel isn't faster, here
	// p1p := pt1Parallel(data).Load()
	// // p1p := pt1Parallel(data)
	// fmt.Println(p1p)
	// if p1p != int64(p1) {
	// 	// if p1p != p1 {
	// 	fmt.Printf("p1: %v\np1p: %v\n", p1, p1p)
	// 	panic("pt1 and pt1Parallel do not produce the same answer")
	// }
	//
	// times := []int64{}
	// nTests := 100_000
	// for i := 0; i < nTests; i++ {
	// 	start := time.Now()
	// 	pt1(data)
	// 	elapsed := time.Since(start).Nanoseconds()
	// 	times = append(times, elapsed)
	// }
	// nonP := meanTime(times)
	// fmt.Printf("Non Parallel: %f avg ns\n", nonP)
	// times = []int64{}
	// for i := 0; i < nTests; i++ {
	// 	start := time.Now()
	// 	pt1Parallel(data)
	// 	elapsed := time.Since(start).Nanoseconds()
	// 	times = append(times, elapsed)
	// }
	// P := meanTime(times)
	// fmt.Printf("Parallel    : %f avg ns\n", P)
	// fmt.Printf("Parallel is  %.4f times faster\n", nonP/P)

	fmt.Println()
	pt2(data)

}

func meanTime(times []int64) float64 {
	var n int64
	for _, t := range times {
		n += t
	}
	return float64(n) / float64(len(times))
}

func getData() []Card {
	b, err := os.ReadFile("../data/4.txt")
	if err != nil {
		panic(err)
	}
	data := string(b)
	cards := []Card{}

	for i, row := range strings.Split(data, "\n") {
		if row == "" {
			continue
		}
		nums := strings.Split(strings.Split(row, ":")[1], "|")
		winningNumbers := []int{}
		for _, num := range strings.Split(strings.TrimSpace(nums[0]), " ") {
			num = strings.TrimSpace(num)
			if num == "" {
				continue
			}
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			winningNumbers = append(winningNumbers, n)
		}

		cardNumbers := []int{}
		for _, num := range strings.Split(strings.TrimSpace(nums[1]), " ") {
			num = strings.TrimSpace(num)
			if num == "" {
				continue
			}

			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			cardNumbers = append(cardNumbers, n)
		}

		slices.Sort(winningNumbers)
		slices.Sort(cardNumbers)
		card := Card{
			id:             i + 1,
			winningNumbers: winningNumbers,
			cardNumbers:    cardNumbers,
		}

		cards = append(cards, card)
	}

	return cards

}
func getTestData() []Card {
	data := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	cards := []Card{}

	for i, row := range strings.Split(data, "\n") {
		if row == "" {
			continue
		}
		nums := strings.Split(strings.Split(row, ":")[1], "|")
		winningNumbers := []int{}
		for _, num := range strings.Split(strings.TrimSpace(nums[0]), " ") {
			num = strings.TrimSpace(num)
			if num == "" {
				continue
			}
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			winningNumbers = append(winningNumbers, n)
		}

		cardNumbers := []int{}
		for _, num := range strings.Split(strings.TrimSpace(nums[1]), " ") {
			num = strings.TrimSpace(num)
			if num == "" {
				continue
			}

			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			cardNumbers = append(cardNumbers, n)
		}

		slices.Sort(winningNumbers)
		slices.Sort(cardNumbers)
		card := Card{
			id:             i + 1,
			winningNumbers: winningNumbers,
			cardNumbers:    cardNumbers,
		}

		cards = append(cards, card)
	}

	return cards

}
