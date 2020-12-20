package main

import . "fmt"

func sumAndSub(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	a := 4
	b := 3
	sum, _ := sumAndSub(a, b)
	Println("sum: ", sum)
}