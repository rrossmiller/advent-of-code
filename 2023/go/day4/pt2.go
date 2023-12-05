package main

import (
	"fmt"
)

func pt2(cards []Card) {
	ans := 0
	cardMap := map[int]int{}

	for _, card := range cards {
		cardMap[card.id]++
		nWinners := 0
		// check if the card has a winning number
		for _, n := range card.cardNumbers {
			if contains(card.winningNumbers, n) {
				nWinners++
				cardMap[card.id+nWinners] += cardMap[card.id] // increment the number of times a card gets copied
			}
		}
	}

	for _, v := range cardMap {
		ans += v
	}

	fmt.Println("pt2:", ans)
}

func p(c []Card) []int {
	n := []int{}
	for _, i := range c {
		n = append(n, i.id)
	}
	return n
}
