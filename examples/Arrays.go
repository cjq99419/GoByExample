package main

import . "fmt"

func main() {

	//常规数组创建
	var a [5]int
	Println("emp: ", a)

	//常用方法，大体类似
	a[1] = 1
	Println("set: ", a)
	Println("get: ", a[1])
	Println("len: ", len(a))

	//一维数组创建（省略var）
	//也可b := [...]int{1, 2, 3, 4, 5}
	//让编译器自己算去吧
	b := [5]int{1, 2, 3, 4, 5}
	Println("dcl: ", b)

	//二维数组，同理
	var c = [2][2]int{{1, 1}, {2, 2}}
	Println("2arr: ", c)
}
