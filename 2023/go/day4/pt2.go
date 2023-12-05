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
				cardMap[card.id+nWinners] += cardMap[card.id]
			}
		}
	}

	for _, v := range cardMap {
		ans += v
	}

	if ans <= 3514 {
		fmt.Println("too low")
	}
	fmt.Println("pt2:", ans)
	// fmt.Println(cardMap)

}

func p(c []Card) []int {
	n := []int{}
	for _, i := range c {
		n = append(n, i.id)
	}
	return n
}
