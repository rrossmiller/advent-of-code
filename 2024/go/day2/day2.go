package day2

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func Run(dat []string) error {
	data, err := intData(dat)
	if err != nil {
		return err
	}

	p1(data)
	p2(data)

	return nil
}

func p1(data [][]int) {
	var wg sync.WaitGroup
	var numOk atomic.Uint32

	// concurrently check every line
	for _, d := range data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ok := checkLine(d)
			if ok {
				numOk.Add(1)
			}
		}()
	}
	wg.Wait()

	fmt.Printf("Part 1: %d\n", numOk.Load())
}

func checkLine(line []int) bool {
	inc := (line[0] - line[1]) < 0
	for i := 1; i < len(line); i++ {
		prev := line[i-1]
		curr := line[i]
		diff := prev - curr

		// Ok if the levels are either all increasing or all decreasing.
		if inc && diff > 0 {
			fmt.Println(inc, diff, "should be neg")
			return false
		} else if !inc && diff < 0 {
			fmt.Println(inc, diff, "should be pos")
			return false
		}

		if inc {
			diff = diff * -1
		}

		// Ok if any two adjacent levels differ by at least one and at most three.
		if diff < 1 || diff > 3 {
			fmt.Println(inc, diff, "wrong diff")
			return false
		}
	}

	return true
}

func p2(data [][]int) {
	// var wg sync.WaitGroup
	// var numOk atomic.Uint32
	var numOk int

	// concurrently check every line
	for _, d := range data {
		// wg.Add(1)
		// go func() {
		// func() {
		fmt.Println(d)
		// defer wg.Done()
		ok := checkLine2(d)
		if ok {
			numOk += 1
		}
		// }()
		fmt.Println()
	}
	// wg.Wait()

	fmt.Printf("Part 2: %d\n", numOk)
	fmt.Println("too low", numOk <= 268)
}

func checkLine2(line []int) bool {
	inc := (line[0] - line[1]) < 0
	for i := 1; i < len(line); i++ {
		prev := line[i-1]
		curr := line[i]
		diff := prev - curr

		// Ok if the levels are either all increasing or all decreasing.
		if inc && diff > 0 {
			fmt.Println(inc, diff, "should be neg")
			fmt.Println([]int{7, 8, 2, 1}, "..", checkLine([]int{7, 8, 2, 1}))
			dat := removeIdx(line, i-1)
			dat1 := removeIdx(line, i)
			fmt.Println(line, dat, i)
			fmt.Println(line, dat1, i)
			return checkLine(dat) || checkLine(dat1)
		} else if !inc && diff < 0 {
			fmt.Println(inc, diff, "should be pos")
			dat := removeIdx(line, i-1)
			dat1 := removeIdx(line, i)
			fmt.Println(line, dat, i)
			fmt.Println(line, dat1, i)
			return checkLine(dat) || checkLine(dat1)
		}

		if inc {
			diff *= -1
		}

		// Ok if any two adjacent levels differ by at least one and at most three.
		if diff < 1 || diff > 3 {
			fmt.Println(inc, diff, "wrong diff")
			dat := removeIdx(line, i-1)
			dat1 := removeIdx(line, i)
			fmt.Println(line, dat, i)
			fmt.Println(line, dat1, i)
			return checkLine(dat) || checkLine(dat1)
		}
	}

	return true
}

func removeIdx(data []int, idx int) []int {
	rtn := make([]int, 0, len(data)-1)
	for i, v := range data {
		if i == idx {
			continue
		}
		rtn = append(rtn, v)
	}
	return rtn
}

//

func intData(data []string) ([][]int, error) {
	rtn := make([][]int, 0, len(data))
	for _, l := range data {
		line := strings.Split(l, " ")
		inner := []int{}
		for _, i := range line {
			v, err := strconv.Atoi(i)
			if err != nil {
				return nil, err
			} else {
				inner = append(inner, v)
			}
		}
		rtn = append(rtn, inner)
	}
	return rtn, nil
}
