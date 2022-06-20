package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		if <-ch != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
