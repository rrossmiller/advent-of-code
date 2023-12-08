package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HandType int8

const (
	HighCard HandType = iota
	Pair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

var handOrderIdx []HandType
var handOrder []HandType

func (h HandType) String() string {
	switch h {
	case HighCard:
		return "HighCard"
	case Pair:
		return "Pair"
	case TwoPair:
		return "TwoPair"
	case ThreeKind:
		return "ThreeKind"
	case FullHouse:
		return "FullHouse"
	case FourKind:
		return "FourKind"
	case FiveKind:
		return "FiveKind"

	}
	return ""
}

var cardOrder []string

type Hand struct {
	cards    string
	rank     int
	bid      int
	handType HandType
}

func (h Hand) String() string {
	return fmt.Sprintf("%v %v %v %v", h.cards, h.rank, h.bid, h.handType)
}

func main() {
	handOrderIdx = []HandType{
		HighCard,
		Pair,
		ThreeKind,
		FourKind,
		FiveKind,
	}
	handOrder = []HandType{
		HighCard,
		Pair,
		TwoPair,
		ThreeKind,
		FullHouse,
		FourKind,
		FiveKind,
	}
	cardOrder = []string{
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"T",
		"J",
		"Q",
		"K",
		"A",
	}

	fmt.Println("*****")
	fmt.Println("Day 7")
	data := getData()
	// data := getTestData()

	p1(data)

}

func getData() []*Hand {
	b, err := os.ReadFile("../data/7.txt")
	if err != nil {
		panic(err)
	}
	data := string(b)
	hands := []*Hand{}

	for _, row := range strings.Split(data, "\n") {
		rowSpl := strings.Split(row, " ")
		if len(rowSpl) == 1 {
			continue
		}
		cards := rowSpl[0]
		bid, err := strconv.Atoi(rowSpl[1])
		if err != nil {
			fmt.Println(">", rowSpl[0], rowSpl[1])
			panic(err)
		}
		h := &Hand{cards: cards, rank: 0, bid: bid, handType: HighCard}
		hands = append(hands, h)
	}

	return hands

}

func getTestData() []*Hand {
	data := `32T3K 765
KK677 28
T55J5 684
KTJJT 220
QQQJA 483`
	// first three are tests
	// 	data := `T55T5 684
	// 555T5 684
	// 55555 684
	// 32T3K 765
	// T55J5 684
	// KK677 28
	// KTJJT 220
	// QQQJA 483`
	hands := []*Hand{}

	for _, row := range strings.Split(data, "\n") {
		rowSpl := strings.Split(row, " ")
		cards := rowSpl[0]
		bid, err := strconv.Atoi(rowSpl[1])
		if err != nil {
			panic(err)
		}
		h := &Hand{cards: cards, rank: 0, bid: bid, handType: HighCard}
		hands = append(hands, h)
	}

	return hands

}
