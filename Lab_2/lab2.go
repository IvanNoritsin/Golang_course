package main

import (
	"fmt"
)

func main() {

	arr := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	var k int
	var max int
	var index int
	max = -1000

	fmt.Printf("Введите число k: ")
	fmt.Scanf("%d", &k)

	for i := 0; i < k; i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
				index = j
			}
		}

		arr[index] = 0
		if i != k-1 {
			max = -1000
		}
	}

	fmt.Printf("Наибольший %d-й элемент массива: %d\n", k, max)

}
