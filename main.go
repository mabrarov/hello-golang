package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, f func(v int)) {
	if t == nil {
		return
	}
	walk(t.Left, f)
	f(t.Value)
	walk(t.Right, f)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t.Left, func(v int) { ch <- v })
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for done := false; !done; {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		done = !ok1 && !ok2
	}
	return true
}

func printTree(t *tree.Tree) {
	walk(t.Left, func(v int) { fmt.Println(v) })
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println("Tree #1:")
	printTree(t1)
	fmt.Println("Tree #2:")
	printTree(t2)
	fmt.Println("Tree #1 is same as #2:", Same(t1, t2))
}
