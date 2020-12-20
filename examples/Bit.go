package main

import . "fmt"

func main() {
	a := 10   //1010
	c := 7    //0111
	Printf("a = %4b\n", a)
	Printf("c = %4b\n", c)
	Printf("^a = %4b\n", ^a)
	Printf("a & c = %4b\n", a & c)
	Printf("a | c = %4b\n", a | c)
	Printf("a &^ c = %4b\n", a &^ c)
	Printf("a >> 1 = %4b\n", a >> 1)
	Printf("a << 1 = %4b\n", a << 1)
}
