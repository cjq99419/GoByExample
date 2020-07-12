package main

import . "fmt"

func main() {

	//if语句大同小异
	a := 0
	if a == 0 {
		Println("略略略")
	} else {
		Println("啦啦啦")
	}

	//注意在if语句中可以进行赋值，且赋值对if和else均生效
	//且Go中的Println不同于java或C++使用+连接或连续<<构造输出流
	//需要把输出按顺序使用，进行分割
	if num := 100; a > -1 {
		Println("num :", num)
	}

}
