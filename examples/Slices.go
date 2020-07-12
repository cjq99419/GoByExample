package main

//slices are a key date type in Go, giving a more powerful interface to sequences than arrays
//slices are typed only by the elements they contain
import . "fmt"

func main() {
	s := make([]string, 3)
	Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	Println("set: ", s)

	s = append(s, "d")
	s = append(s, "e")
	Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	Println("cpy: ", c)

	s1 := s[2:3]
	Println("sl1: ", s1)

	s2 := s1[:cap(s1)-1]
	Println("sl2: ", s2)

	d := []string{"aa", "bb", "cc"}
	Println("dcl: ", d)
}
