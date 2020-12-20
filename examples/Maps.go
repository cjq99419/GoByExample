package main

import . "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	Println("map:", m)

	v1 := m["k1"]
	Println("v1: ", v1)

	Println("len:", len(m))

	delete(m, "k1")
	Println("map:", m)

	aaa, prs := m["k2"]
	Println("prs:", prs)
	Println("aaa:", aaa)

	n := map[string]int{"foo": 1, "bar": 2}
	Println("map:", n)
}