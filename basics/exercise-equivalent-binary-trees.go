package main

import "golang.org/x/tour/tree"

// Walk through the tree t sending all values from the tree to the chanel ch.
func Walk(t *tree.Tree, ch chan int)

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool

func main() {
}
