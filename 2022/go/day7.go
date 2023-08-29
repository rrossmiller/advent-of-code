package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	parent   *Dir
	children []*Dir
	path     string
	size     uint64
}

func (d Dir) String() string {
	par := "nil"
	if d.parent != nil {
		par = d.parent.path
	}
	return fmt.Sprintf("Dir {%v %v %v %v}", par, d.path, d.size, len(d.children))
}

func Seven(maxSize uint64) {
	dat := read7()

	var parent *Dir
	lines := strings.Split(dat, "\n")[1:]
	root := Dir{parent: parent, path: "/", size: 0}
	current := &root

	for _, line := range lines {
		elems := strings.Split(line, " ")

		// if the line is a command
		if elems[0] == "$" {
			// if the command is change dir
			if elems[1] == "cd" {
				// cd .. means current dir becomes the parent of the current dir
				if elems[2] == ".." {
					current = parent
					parent = parent.parent

				} else { // cd into the new dir
					parent = current
					current = &Dir{parent: parent, path: elems[2], size: 0}
					parent.children = append(parent.children, current)
				}
			}
		} else if i, err := strconv.Atoi(elems[0]); err == nil {
			current.size += uint64(i)

			for par := current.parent; par != nil; par = par.parent {
				par.size += uint64(i)

			}

		}

	}
	// fmt.Println(current)
	// fmt.Println(current.size == (206810 + 20178))
	// fmt.Println(root)

	// search for dirs with size <= maxSize
	//DFS
	var sum uint64
	stack := []*Dir{}
	stack = append(stack, &root)

	// while stack is not empty:
	for len(stack) > 0 {
		//     current = pop from stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//     process current node
		if current.size <= maxSize {
			sum += current.size

		}
		//     for each child in current's children:
		for _, c := range current.children {
			//         push child onto stack
			stack = append(stack, c)
		}
	}
	fmt.Println(">>>", sum)

}

func read7() string {
	fBytes, err := os.ReadFile("../data/7.txt")
	Check(err)

	return string(fBytes)
}
