package main

import . "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		Println("key:", k)
	}

	for i, c := range "go" {
		Println(i, c)
	}
}
