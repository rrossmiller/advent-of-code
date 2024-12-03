package day1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Run(data []string) error {

	// get the left and right numbers
	lIds := make([]int, 0, len(data))
	rIds := make([]int, 0, len(data))
	for _, line := range data {
		if len(line) == 0 {
			continue
		}
		nums := strings.Split(line, " ")

		l, err := strconv.Atoi(nums[0])
		if err != nil {
			return err
		}
		r, err := strconv.Atoi(nums[len(nums)-1])
		if err != nil {
			return err
		}
		lIds = append(lIds, l)
		rIds = append(rIds, r)
	}
	// lists in ascending order
	slices.Sort(lIds)
	slices.Sort(rIds)
	p1(lIds, rIds)
	p2(lIds, rIds)
	return nil
}

func p1(lIds, rIds []int) {
	dist := 0
	for i := range len(lIds) {
		l := lIds[i]
		r := rIds[i]

		d := l - r
		if d < 0 {
			dist += d * -1
		} else {
			dist += d
		}
	}
	fmt.Println("Part 1:", dist)

}

type counter map[int]int

func (c counter) countN(k, v int) {
	_, prs := c[k]
	if prs {
		c[k] += v
	} else {
		c[k] = v
	}

}
func (c counter) count(k int) {
	c.countN(k, 1)
}
func p2(lIds, rIds []int) {
	rCount := counter{}
	for i := range len(lIds) {
		rCount.count(rIds[i])
	}

	dist := 0
	for _, k := range lIds {
		rv := rCount[k]
		dist += k * rv
	}
	fmt.Println("Part 2:", dist)
}
