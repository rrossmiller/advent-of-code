package main

import (
	"fmt"
	"os"
	"strings"
)

func Day12() {
	path := "test.txt"
	data := getData(path)
	fmt.Println(data)
}

func getData12(path string) [][]string {
	fBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(fBytes), "\n")
	rtn := [][]string{}

	for _, r := range rows {
		rtn = append(rtn, strings.Split(r, ""))
	}
	return rtn
}
