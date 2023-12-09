package main

import (
	// "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := getData()
	// data = getTestData()
	// for _, n := range data {
	// 	fmt.Printf("%T: %v\n", n[0], n)
	// }

	p1(data,true)
	p1(data,false)
}

func getData() [][]int64 {
	b, err := os.ReadFile("../data/9.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(b), "\n")
	return process(data)

}
func getTestData() [][]int64 {
	f := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	data := strings.Split(f, "\n")
	return process(data)
}
func process(data []string) [][]int64 {
	rtn := [][]int64{}
	for _, l := range data {
		if len(l) == 0 {
			continue
		}
		inner := []int64{}
		for _, n := range strings.Split(l, " ") {
			val, err := strconv.ParseInt(n, 10, 64)
			if err != nil {
				panic(err)
			}
			inner = append(inner, val)
		}
		rtn = append(rtn, inner)
	}
	return rtn

}
