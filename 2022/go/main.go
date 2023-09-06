package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must pass in the number of the day to run (i.e., '1')")
		os.Exit(1)
	}

	if os.Args[1] == "1" {
		One()
	} else if os.Args[1] == "7" {
		Seven(100_000)
	} else if os.Args[1] == "8" {
		Eight()
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
