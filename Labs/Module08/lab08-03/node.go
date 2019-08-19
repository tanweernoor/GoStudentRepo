// Liked List code for lab 08-03
package main

import "fmt"

type node struct {
	prev     *node
	contents int
	next     *node
}
type list struct {
	start *node
	end   *node
}

func makeEmptyList() *list {
	// Just makes an empty list.
	return &list{nil, nil}
}

func addNode(f int, l *list) {
	// Given a list and an int, the node is
	// created and added to the end of the list
	node := &node{nil, f, nil}
	if l.end == nil {
		// special case if the list is emty
		l.end, l.start = node, node
		return
	}
	// Add and update the other nodes
	node.prev, l.end.next = l.end, node
	l.end = node
	return
}

func printList(l *list) {
	if l.start == nil {
		fmt.Println("Empty List")
		return
	}
	for n := l.start; n != nil; n = n.next {
		fmt.Println("Node:", n.contents)
	}
}
