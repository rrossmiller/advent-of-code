package main

import (
	"fmt"
)

// todo custom bin search
// paraellize
func pt1(data []Card) {
	ans := 0
	for _, d := range data {
		fmt.Println(d)
		cardVal := 0
		for _, n := range d.cardNumbers {
			for _, w := range d.winningNumbers {
				if n == w {
					if cardVal == 0 {
						cardVal++
					} else {
						cardVal *= 2
					}
				}
			}
		}
		ans += cardVal
	}
	fmt.Println("pt1:", ans)
}
