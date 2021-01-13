package main

import "fmt"

type node struct {
	Value int
	Previous *node
	Next *node
}

func addnode1(t *node, v int) int {
	if root1 == nil {
		t = &node{v, nil, nil}
		root1 = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1

	}

	if t.Next == nil {
		temp := t
		t.Next = &node{v, temp, nil}
		return -2
	}
	return addnode1(t.Next, v)
}

func traverse1(t *node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func reverse1(t *node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}
	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d ->", temp.Value)
	fmt.Println()
}

func size1(t *node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupnode1(t *node, v int) bool {
	if root1 == nil {
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupnode1(t.Next, v)
}

var root1 = new(node)

func main() {
	fmt.Println(root1)
	root1 = nil
	traverse1(root1)
	addnode1(root1, 1)
	addnode1(root1, 1)
	traverse1(root1)
	addnode1(root1, 10)
	addnode1(root1, 5)
	addnode1(root1, 0)
	addnode1(root1, 0)
	traverse1(root1)
	addnode1(root1, 100)
	fmt.Println("Size : ", size1(root1))
	traverse1(root1)
	reverse1(root1)
}