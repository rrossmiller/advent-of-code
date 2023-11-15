package days

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Point struct {
	X int
	Y int
}
type heuristic func(Point, Point) float64

func Day12() {
	path := "test.txt"
	data, endPos := getData12(path)
	fmt.Println(data)
	fmt.Println(endPos)
	fmt.Println()

	ans := aStar(data, endPos, dist)
	if strings.Contains(path, "test") {
		fmt.Println("Ans Correct:", ans == 31)
	}
}

func dist(p, e Point) float64 {
	x := math.Pow(float64(p.X-e.X), 2)
	y := math.Pow(float64(p.Y-e.Y), 2)
	return math.Sqrt(x + y)
}

func aStar(data [][]string, endPos Point, h heuristic) int {
	steps := 0
	start := Point{0, 0}

	// The set of discovered nodes that may need to be (re-)expanded.
	// Initially, only the start node is known.
	// This is usually implemented as a min-heap or priority queue rather than a hash-set.
	openSet := []Point{start}

	// For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from the start
	// to n currently known.
	// cameFrom := make(map[Point]Point) //an empty map

	// For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
	gScore := make(map[Point]float64) //map with default value of Infinity
	// gScore[start] := 0
	fmt.Println(gScore[start])

	// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
	// how cheap a path could be from start to finish if it goes through n.
	fScore := make(map[Point]float64) //map with default value of Infinity
	fScore[start] = h(start, endPos)

	// while openSet is not empty
	for len(openSet) > 0 {

		//        // This operation can occur in O(Log(N)) time if openSet is a min-heap or a priority queue
		//        current := the node in openSet having the lowest fScore[] value
		//        if current = goal
		//
		//        openSet.Remove(current)
		//        for each neighbor of current
		//            // d(current,neighbor) is the weight of the edge from current to neighbor
		//            // tentative_gScore is the distance from start to the neighbor through current
		//            tentative_gScore := gScore[current] + d(current, neighbor)
		//            if tentative_gScore < gScore[neighbor]
		//                // This path to neighbor is better than any previous one. Record it!
		//                cameFrom[neighbor] := current
		//                gScore[neighbor] := tentative_gScore
		//                fScore[neighbor] := tentative_gScore + h(neighbor)
		//                if neighbor not in openSet
		//                    openSet.add(neighbor)
		//
		break
	}
	//    // Open set is empty but goal was never reached
	//    // return failure
	return steps
}

func AddPoints(points []Point, p Point) []Point {
	if !slices.Contains(points, p) {
		points = append(points, p)
		return points
	}
	return points
}
func getData12(path string) ([][]string, Point) {
	fBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(fBytes), "\n")
	rtn := [][]string{}
	var endPos Point

	for i, r := range rows {
		rtn = append(rtn, strings.Split(r, ""))

		for j, c := range rtn[i] {
			if c == "E" {
				endPos.X = i
				endPos.Y = j
			}
		}
	}

	return rtn, endPos
}
