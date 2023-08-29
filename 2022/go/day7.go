package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	parent *Dir
	path   string
	size   int
}

func (d Dir) String() string {
	return fmt.Sprintf("Dir {%v %v %v}", d.parent.path, d.path, d.size)
}

func Seven(maxSize int) {
	dat := read7()
	// go through each line
	var parent *Dir
	var current Dir
	// current := Dir{parent: parent, path: "/", size: 0}

	for _, line := range strings.Split(dat, "\n") {
		elems := strings.Split(line, " ")

		// if the line is a command
		if elems[0] == "$" {
			// if the command is change dir
			if elems[1] == "cd" {
				// cd .. means current dir becomes the parent of the current dir
				if elems[2] == ".." {
					current = *parent
					parent = parent.parent

				} else { // cd into the new dir
					parent = &Dir{parent: current.parent, path: current.path, size: current.size}
					current = Dir{parent: parent, path: elems[2], size: 0}
				}
			}
		} else if i, err := strconv.Atoi(elems[0]); err == nil {
			current.size += i

			x := 0
			for par := current.parent; par != nil; par = par.parent {
				par.size += i
				x += 1

			}

		}

	}

	// loop lines done
    // search for dirs with size <= maxSize 
    // sum of dir sizes

}

func read7() string {
	fBytes, err := os.ReadFile("../data/7.txt")
	Check(err)

	return string(fBytes)
}
