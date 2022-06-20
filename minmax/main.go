package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(Min(2, 3))
	fmt.Println(Max("23", "2"))
}
