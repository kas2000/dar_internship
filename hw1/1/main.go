package main

import (
	"fmt"
	"time"
)


func main() {
	var arr1 = []int{63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var arr2 = []int{63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var arr3 = []int{63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var arr4 = []int{63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 63, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	test(arr1)
	test(arr2)
	test(arr3)
	test(arr4)


}

func test(nums []int) {
	t0 := time.Now()
	bubbleSort(nums)
	t1 := time.Now()

	fmt.Println(nums)
	fmt.Println(t1.Sub(t0))
}

/* 
	O(n)
	Average case = n^2
	Worst case = n^2
	Best case = n
*/
func bubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}