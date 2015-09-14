package main

import "golang.org/x/tour/tree"
import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, head bool) {
	if t == nil {
		return
	}
	Walk(t.Left, ch, false)
	ch <- t.Value
	Walk(t.Right, ch, false)
	if head {
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1, true)
	go Walk(t2, ch2, true)
	for i := range ch1 {
		j := <-ch2
		if i != j {
			return false
		}
	}
	return true
}

func main() {
	same1 := Same(tree.New(1), tree.New(1))
	same2 := Same(tree.New(2), tree.New(1))
	fmt.Printf("same1 ? %v\n", same1)
	fmt.Printf("same2 ? %v\n", same2)
}
