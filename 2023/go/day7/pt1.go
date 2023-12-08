package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func p1(hands []*Hand) {
	// figure out what hand type each hand is
	for _, h := range hands {
		counter := map[string]int{}
		for _, c := range strings.Split(h.cards, "") {
			counter[c]++
		}

		// list of counts to start to determine the hand
		counts := []int{}
		for _, v := range counter {
			counts = append(counts, v)
		}

		slices.Sort(counts)
		slices.Reverse(counts)

		if counts[0] == 2 && counts[1] == 2 {
			h.handType = TwoPair
		} else if counts[0] == 3 && counts[1] == 2 {
			h.handType = FullHouse
		} else {
			h.handType = handOrderIdx[counts[0]-1]
		}
	}

	// group by hand type
	groups := map[HandType][]*Hand{}
	for _, h := range hands {
		groups[h.handType] = append(groups[h.handType], h)
	}

	// set rank
	rank := len(hands)
	for i := len(handOrder) - 1; i >= 0; i-- {
		handsGroup := groups[handOrder[i]]
		if len(handsGroup) == 1 {
			handsGroup[0].rank = rank
			rank--
		} else {
			// sort by the cards in the hand
			slices.SortFunc(handsGroup, func(a, b *Hand) int {
				aRunes := strings.Split(a.cards, "")
				bRunes := strings.Split(b.cards, "")
				for i := 0; i < len(a.cards); i++ {
					r := slices.Index(cardOrder, aRunes[i]) - slices.Index(cardOrder, bRunes[i])
					if r < 0 {
						return 1
					} else if r > 0 {
						return -1
					}
				}
				return 0
			})
			for _, h := range handsGroup {
				h.rank = rank
				rank--
			}

		}

	}

	// order by rank
	slices.SortFunc(hands, func(a, b *Hand) int {
		return cmp.Compare(a.rank, b.rank)
	})

	// results
	fmt.Println(".......")
	sum := 0
	for _, h := range hands {
		sum += h.rank * h.bid
	}
	fmt.Println("pt1:", sum)

}
