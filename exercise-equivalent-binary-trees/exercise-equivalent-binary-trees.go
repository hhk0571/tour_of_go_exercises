/*
https://tour.golang.org/concurrency/7
https://tour.golang.org/concurrency/8

Task: implement the Walk() function to walk through the binary tree

*/

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if t != nil {
		_walk(t.Left, ch)
		fmt.Println("sent", t.Value)
		ch <- t.Value
		_walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		if v2 := <-ch2; v2 != v1 {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for i := range ch {
		fmt.Println("received", i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))

}

