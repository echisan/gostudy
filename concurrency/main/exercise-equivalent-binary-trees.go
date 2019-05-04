package main

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func Walk(t *Tree, ch chan int) {
	walkTree(t, ch)
}

func walkTree(t *Tree, ch chan int) {
	if t.Left != nil {
		walkTree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkTree(t.Right, ch)
	}
}

func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
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

}
