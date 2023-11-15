package days

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func One() {
	dataString := unzip()
	data := strings.Split(dataString, "\n")
	// ptOne(data)
	ptTwo(data)
}
func ptTwo(data []string) {
	// get the elf with the highest cal stored
	start := time.Now()
	max3 := [3]int{}
	runningTtl := 0
	for _, v := range data {
		if v == "" {
			max3 = Swap(max3, runningTtl)
			runningTtl = 0
		} else {
			i, _ := strconv.Atoi(v)
			runningTtl += i
		}
	}

	// sum the top 3
	max := 0
	for _, v := range max3 {
		max += v
	}

	runtime := time.Since(start)
	fmt.Println(runtime)
	fmt.Println("top 3 Cals: ", max)
}

func Swap(max3 [3]int, ttl int) [3]int {
	for i := 0; i < 3; i++ {
		if ttl > max3[i] {
			tmp := max3[i]
			max3[i] = ttl
			ttl = tmp
		}
	}
	return max3
}

func ptOne(data []string) {
	// get the elf with the highest cal stored
	start := time.Now()
	max := 0
	runningTtl := 0
	for _, v := range data {
		if v == "" {
			if runningTtl > max {
				max = runningTtl
			}
			runningTtl = 0
		} else {
			i, _ := strconv.Atoi(v)
			runningTtl += i
		}
	}

	runtime := time.Since(start)
	fmt.Println(runtime)

	fmt.Println("highest Cals: ", max)
}

func unzip() string {
	dataPath := "../data/1.zip"
	_, err := os.Stat(dataPath)
	Check(err)

	// open zip archive for reading
	r, err := zip.OpenReader(dataPath)
	Check(err)
	defer r.Close()

	// Iterate through the files in the archive,
	b := []byte{}
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		Check(err)
		contents, err := io.ReadAll(rc)
		Check(err)
		b = append(b, contents...)
		rc.Close()
	}

	// b, _ := os.ReadFile("../data/1.txt")

	return string(b)
}
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
