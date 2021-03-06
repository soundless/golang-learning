package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk through the tree t sending all values from the tree to the chanel ch.
func Walk(t *tree.Tree, ch chan int) {
	traverse(t, ch)
	close(ch)
}

func traverse(t *tree.Tree, ch chan int) {
	if t != nil {
		traverse(t.Left, ch)
		ch <- t.Value
		traverse(t.Right, ch)
	}
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
