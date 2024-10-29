package main

import (
	"os"
	"strings"
)

func main() {
	test := false

	var data string
	if test {
		data = testData()
	} else {
		data = getData()
	}
	lines := strings.Split(data, "\n")
	p1(lines)
}

func testData() string {
	dat := `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

	return dat
}
func getData() string {
	b, err := os.ReadFile("../data/14.txt")
	if err != nil {
		panic(err)
	}
	data := string(b)

	return data

}
