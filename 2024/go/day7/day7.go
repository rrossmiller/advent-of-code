package day7

import (
	"fmt"
)

func Run(data []string) error {
	for _, l := range data {
		fmt.Println(l)
	}

	err := pt1()
	if err != nil {
		return err
	}
	return nil
}

func pt1() error {
	return nil
}
