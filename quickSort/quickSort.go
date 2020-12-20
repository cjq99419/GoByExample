package main

import "fmt"

func main() {
	arr := []int{1, 5, 2, 7, 4, 3}
	QuickSort(arr)
	fmt.Println(arr)
}

func QuickSort(intList []int) {
	if len(intList) <= 1 {
		return
	}
	flag := intList[0]
	left, right := 0, len(intList)-1

	for i := 1; i <= right; {
		if intList[i] > flag {
			intList[i], intList[right] = intList[right], intList[i]
			right--
		} else {
			intList[i], intList[left] = intList[left], intList[i]
			i++
			left++
		}
	}
	// 递归
	QuickSort(intList[:left])
	QuickSort(intList[left+1:])
}
