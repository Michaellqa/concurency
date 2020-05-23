package gotour

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, root bool) {
	if t == nil {
		return
	}
	Walk(t.Left, ch, false)
	ch <- t.Value
	Walk(t.Right, ch, false)

	if root {
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1, true)
	go Walk(t2, ch2, true)

	// all nodes from tree1 the same as in the tree2
	var difFound bool
	for v1 := range ch1 {
		v2, closed := <-ch2
		if !closed || v1 != v2 {
			difFound = true
		}
	}

	result := !difFound

	// if there more values left in tree2 after tree1 has been walked
	for range ch2 {
		result = false
	}

	return result
}
