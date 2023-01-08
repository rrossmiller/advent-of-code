package main

func main() {
	One()
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
