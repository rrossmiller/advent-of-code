package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Monkey struct {
	items       []int
	op          string
	opVal       int // if this is math.MinInt, use old
	testVal     int
	passTo      []int
	inspections int64
}

func Day11() {
	path := "../data/day_11.txt"
	// path := "../python/test.txt"
	monkeys := getData(path)
	for i, m := range monkeys {
		fmt.Printf("Monkey %d: %v\n", i, m)
	}
	fmt.Println()

	x := 10_000
	for n := 0; n < x; n++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				i := m.items[0]
				m.items = m.items[1:]
				m.testAndThrow(i, monkeys)
			}
		}
	}

	fmt.Println()
	ans := []int64{}
	for _, m := range monkeys {
		ans = append(ans, m.inspections)
	}
	slices.Sort(ans)
	slices.Reverse(ans)

	fmt.Println(ans)
	s := ans[0] * ans[1]
	fmt.Println(s)

	if strings.Contains(path, "test") {
		fmt.Println(s == 2713310158) //test ans
		fmt.Println(s > 2713310158)
		fmt.Println(s < 2713310158)
	} else {

		if x <= 207129660 {
			fmt.Println("too low")
		} else if x >= 16858797605 {
			fmt.Println("too high")
		}
	}

}

func (m Monkey) operate(a int) int {
	opVal := m.opVal
	if m.opVal == math.MinInt {
		opVal = a
	}

	switch m.op {
	case "+":
		return a + opVal
	case "-":
		return a - opVal
	case "*":
		return a * opVal
	case "/":
		return a / opVal
	}
	return 0
}

func (m *Monkey) testAndThrow(a int, monkeys []*Monkey) {
	a = m.operate(a) // / 3

	m.inspections++
	if a%m.testVal == 0 {
		monkeys[m.passTo[0]].items = append(monkeys[m.passTo[0]].items, a)
	} else {
		monkeys[m.passTo[1]].items = append(monkeys[m.passTo[1]].items, a)
	}
}

func getData(path string) []*Monkey {
	fBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	x := [][]string{}
	for n, i := range strings.Split(string(fBytes), "\n\n") {
		x = append(x, []string{})

		for _, j := range strings.Split(i, "\n")[1:] {
			j = strings.TrimSpace(j)
			x[n] = append(x[n], j)
		}
	}
	y := [][]string{}
	for n, i := range x {
		y = append(y, []string{})
		for _, j := range i {
			if utf8.RuneCountInString(j) > 0 {
				y[n] = append(y[n], strings.Split(j, ":")[1])
			}
		}
	}

	monkeys := []*Monkey{}
	for _, m := range y {
		// Items
		items := []int{}
		for _, i := range strings.Split(m[0], ",") {
			i = strings.TrimSpace(i)
			num, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			items = append(items, num)
		}
		// fmt.Println("items:", items)

		// Operation
		operation := strings.Split(m[1], " ")
		op := operation[len(operation)-2]
		opValStr := operation[len(operation)-1]
		isOld := opValStr == "old"
		opVal := math.MinInt
		if !isOld {
			opVal, err = strconv.Atoi(opValStr)
			if err != nil {
				panic(err)
			}
		}
		// if isOld {
		// 	// fmt.Println("op, opVal:", op, opValStr)
		// } else {
		// 	// fmt.Println("op, opVal:", op, opVal)
		// }

		// Test
		spl := strings.Split(m[2], " ")
		testVal, err := strconv.Atoi(spl[len(spl)-1])
		if err != nil {
			panic(err)
		}
		spl = strings.Split(m[3], " ")
		passT, err := strconv.Atoi(spl[len(spl)-1])
		if err != nil {
			panic(err)
		}
		spl = strings.Split(m[4], " ")
		passF, err := strconv.Atoi(spl[len(spl)-1])
		passTo := []int{passT, passF}
		// fmt.Println("test:", testVal)
		// fmt.Println("pass:", passTo)

		// init monkey and append
		inspections := int64(0)
		monk := &Monkey{items, op, opVal, testVal, passTo, inspections}
		// fmt.Println("Monkey", monk)
		monkeys = append(monkeys, monk)
		// fmt.Println()
	}

	return monkeys
}
