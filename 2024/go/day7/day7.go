package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type row struct {
	Ans      int
	Operands []int
}

func Run(data []string) error {
	rows := make([]row, len(data))
	for i, l := range data {
		// get the ans
		spl := strings.Split(l, ": ")
		ans, err := strconv.Atoi(spl[0])
		if err != nil {
			return err
		}
		// get the operands
		opsSpl := strings.Split(spl[1], " ")
		ops := make([]int, len(opsSpl))
		for j, o := range opsSpl {
			n, err := strconv.Atoi(o)
			if err != nil {
				return err
			}
			ops[j] = n
		}
		rows[i] = row{ans, ops}
	}

	// rows = append(rows, rows...)
	// rows = append(rows, rows...)
	// rows = append(rows, rows...)
	// rows = append(rows, rows...)
	// rows = append(rows, rows...)
	// rows = append(rows, rows...)
	// rows = append(rows, rows...)
	// rows = append(rows, rows...)
	// rows = append(rows, rows...)
	// rows = append(rows, rows...)
	fmt.Println(len(rows))
	err := pt1(rows)
	if err != nil {
		return err
	}
	return nil
}

func pt1(rows []row) error {
	// fmt.Println(time.Now())
	p := message.NewPrinter(language.English)
	var wg sync.WaitGroup
	var ans atomic.Uint64
	var total atomic.Uint64 // track the number of combos of inputs ran
	bar := progressbar.Default(int64(len(rows)))

	// find the answer to every line
	// nWorkers := runtime.NumCPU() * 2
	// nWorkers := len(rows)
	nWorkers := min(100, len(rows))

	work := make(chan row, len(rows))
	for range nWorkers {
		// for _, r := range rows {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := range work {
				if findAns(r, &total) {
					ans.Add(uint64(r.Ans))
				}
				bar.Add(1)
				bar.Describe(p.Sprintf("%d ran (ans: %d)", total.Load(), ans.Load()))
			}
		}()
	}

	for _, r := range rows {
		work <- r
	}

	close(work)
	wg.Wait()
	bar.Close()

	fmt.Println()
	p.Printf("Total %d ran (ans: %d)\n", total.Load(), ans.Load())
	fmt.Printf("Pt1: %d\n", ans.Load())
	if ans.Load() <= 107588740373 {
		fmt.Println("too low")
	}
	return nil
}

func findAns(r row, total *atomic.Uint64) bool {
	// how many different combinations of + and * could there be
	numOps := int(math.Pow(2, float64(len(r.Operands))-1))
	operators := generateCombinations([]string{"+", "*"}, numOps)
	// fmt.Println(operators)

	for _, combo := range operators {
		ans := r.Operands[0]
		total.Add(1) // track the number of combos of inputs ran
		for i := 1; i < len(r.Operands); i++ {
			op := combo[i-1]
			switch op {
			case "+":
				ans = ans + r.Operands[i]
			case "*":
				ans = ans * r.Operands[i]
			}
			if ans > r.Ans {
				break
			}
		}

		if ans == r.Ans {
			// fmt.Println(ans)
			return true
		}
	}

	return false
}

// Function to generate combinations
// thank you chatgpt
func generateCombinations(elements []string, length int) [][]string {
	// Resulting slice to store combinations
	var result [][]string

	// Helper function to generate combinations recursively
	var helper func([]string, int)

	// Helper function to handle the recursion
	helper = func(current []string, depth int) {
		// If current slice reached the desired length, add it to the result
		if depth == length {
			// Append the current combination to the result
			combination := make([]string, len(current))
			copy(combination, current)
			result = append(result, combination)
			return
		}

		// Recursively fill the combination at the current depth
		for _, elem := range elements {
			// Add the element to the current combination
			helper(append(current, elem), depth+1)
		}
	}

	// Start the recursive helper with an empty slice and depth 0
	helper([]string{}, 0)
	return result
}
