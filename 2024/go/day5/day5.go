package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Run(data []string) error {
	ordering := map[int][]int{}
	updates := [][]int{}
	a := true
	for _, l := range data {
		if l == "" {
			a = false
			continue
		}
		if a {
			nums, err := parseLine(l, "|")
			if err != nil {
				return err
			}
			_, prs := ordering[nums[0]]
			if prs {
				ordering[nums[0]] = append(ordering[nums[0]], nums[1])
			} else {
				ordering[nums[0]] = []int{nums[1]}
			}
			// ordering = append(ordering, order{nums[0], nums[1]})
		} else {
			nums, err := parseLine(l, ",")
			if err != nil {
				return err
			}
			updates = append(updates, nums)
		}

	}

	err := p1(ordering, updates)
	if err != nil {
		return err
	}

	err = p2(ordering, updates)
	if err != nil {
		return err
	}
	return nil
}

func p1(ordering map[int][]int, updates [][]int) error {
	ans := 0
	// for every update
	for _, u := range updates {
		ok := true
		// don't need to check first index (nothing can come before it)
		for i, k := range u[1:] {
			// check ordering rules for k
			for _, n := range ordering[k] {
				// no n should be before k in the slice
				for j := i; j >= 0; j-- {
					if u[j] == n {
						ok = false
						break
					}
				}
			}
		}
		if ok {
			mid := len(u) / 2
			ans += u[mid]
		}
	}

	fmt.Printf("Part 1: %d\n", ans)
	return nil
}
func p2(ordering map[int][]int, updates [][]int) error {
	ans := 0
	// for every update
	for _, u := range updates {
		ok := true
		// don't need to check first index (nothing can come before it)
		for i, k := range u[1:] {
			// check ordering rules for k
			for _, n := range ordering[k] {
				// no n should be before k in the slice
				for j := i; j >= 0; j-- {
					if u[j] == n {
						ok = false
						break
					}
				}
			}
		}
		if !ok {
			slices.SortFunc(u, func(a, b int) int {
				// -1 if a before b
				if slices.Contains(ordering[a], b) {
					return -1
				} else if slices.Contains(ordering[b], a) {
					return 1
				}
				return 0
			})
			mid := len(u) / 2
			ans += u[mid]
		}
	}

	fmt.Printf("Part 2: %d\n", ans)
	return nil
}
func parseLine(line, sp string) ([]int, error) {
	spl := strings.Split(line, sp)
	rtn := make([]int, len(spl))
	for i, s := range spl {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		rtn[i] = n

	}

	return rtn, nil
}
