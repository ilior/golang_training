package main

import "fmt"

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func MergeSort(data []int, c chan []int) {
	if len(data) <= 1 {
		c <- data
		return
	}
	mid := len(data) / 2
	l, r := make(chan []int), make(chan []int)
	go MergeSort(data[:mid], l)
	go MergeSort(data[mid:], r)
	c <- Merge(<-l, <-r)
}

func main() {
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}

	c := make(chan []int)

	go MergeSort(data, c)
	fmt.Printf("%v\n%v\n", <-c)
}
