package main

import . "fmt"

func sum(nums ...int) {
	Println("nums: ", nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	Println("total: ", total)
}

func main() {
	sum(1, 2, 3)

	a := []int{4, 5, 6, 7}
	sum(a...)
}
