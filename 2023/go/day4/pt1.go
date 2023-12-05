package main

import "fmt"

// todo custom bin search
// paraellize
func pt1(cards []Card) {
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
	fmt.Println("pt1:", ans)
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
