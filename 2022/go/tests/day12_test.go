package main_test

import (
	"aoc/days"
	"fmt"
	"testing"
)

func TestAddPoint(t *testing.T) {
	p := days.Point{X: 0, Y: 0}
	points := []days.Point{p}

	points = days.AddPoints(points, p)
	if len(points) > 1 {
		t.Error("should not add a point")
	}
	fmt.Println(points)

	p = days.Point{X: 1, Y: 0}
	points = days.AddPoints(points, p)
	if len(points) != 2 {
		t.Error("Point should have been added")
	}
	fmt.Println(points)
}
