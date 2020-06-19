package main

import (
	. "fmt"
	"math"
)

const s string = "constant"

func main() {
	Println(s)

	const n = 500000000

	const d = 3e20 / n
	Println(d)

	Println(int64(d))

	Println(math.Sin(n))
}
