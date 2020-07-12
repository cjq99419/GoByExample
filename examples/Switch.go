package main

import . "fmt"
import "time"

func main() {

	//基础switch，与其他语言类似
	a := 5
	switch a {
	case 1:
		Println("the value of a is one")
	default:
		Println("the value of a is not one")
	}

	//原文：you can use commas to separate multiple expressions in the sa,e case statement
	//即你可以在一个case里写多个表达式，即在一个case中进行多项匹配
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: //当周六周日时输出
		Println("it is the weekend")
	default:
		Println("it is a weekday")
	}

	//原文：switch without an expression is an alternate way to express if/else logic
	//即switch在未指定表达式时可以当作if/else来使用
	switch {
	case a < 0:
		Println("a is less than 0")
	default:
		Println("a is not less than 0")
	}

	//interface不做过多解释
	//此处进行一个类型判断，类似枚举
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			Println(i, " is a bool")
		case int:
			Println(i, " is an int")
		case string:
			Println(i, " is a string")
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("aaa")
}
