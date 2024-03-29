package main

import (
	"sync"
	"sync/atomic"
)

func pt1(cards []Card) int {
	ans := 0
	for _, card := range cards {
		cardVal := 0
		for _, n := range card.cardNumbers {
			if contains(card.winningNumbers, n) {
				if cardVal == 0 {
					cardVal++
				} else {
					cardVal *= 2
				}
			}

		}
		ans += cardVal
	}
	return ans
}

// // this is almost 2x slower than synchronous
// func pt1Parallel(cards []Card) int{
// 	ans := 0
// 	nWorkers := 10
// 	var wg sync.WaitGroup
// 	wg.Add(nWorkers)
// 	jobs := make(chan Card, len(cards))
// 	results := make(chan int, len(cards))
// 	for i := 0; i < nWorkers; i++ {
// 		wg.Add(1)
// 		go func() {
// 			run(jobs, results)
// 			wg.Done()
// 		}()
// 	}
//
// 	for _, card := range cards {
// 		jobs <- card
// 	}
//
// 	close(jobs)
// 	wg.Wait()
// 	close(results)
//
// 	for r := range results {
// 		ans += r
// 	}
// 	return ans
// }

func pt1Parallel(cards []Card) *atomic.Int64 {
	var ans atomic.Int64

	nWorkers := 2
	var wg sync.WaitGroup
	jobs := make(chan Card, len(cards))
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go func() {
			run(jobs, &ans)
			wg.Done()
		}()
	}

	for _, card := range cards {
		jobs <- card
	}

	close(jobs)
	wg.Wait()
	return &ans
}

func run(jobs <-chan Card, ans *atomic.Int64) {
	for card := range jobs {
		cardVal := 0
		for _, n := range card.cardNumbers {
			if contains(card.winningNumbers, n) {
				if cardVal == 0 {
					cardVal++
				} else {
					cardVal *= 2
				}
			}
		}
		ans.Add(int64(cardVal))
	}
}

// binary search
func contains(s []int, i int) bool {
	l, r := 0, len(s)-1
	for l <= r {
		idx := (l + r) / 2
		v := s[idx]
		if v == i {
			return true
		} else if i < v {
			r = idx - 1
		} else {
			//i > v {
			l = idx + 1
		}
	}
	return false
}
