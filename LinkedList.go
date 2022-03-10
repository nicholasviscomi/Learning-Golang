package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	value int
	next *Node
}

type List struct {
	head, tail *Node
}

func (l *List) Append(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	l.tail.next = n
	l.tail = n
}

func (l *List) Search(val int) *Node {
	curr := l.head
	for curr.next != nil {
		if curr.value == val {
			return curr
		} else {
			curr = curr.next
		}
	}
	if curr.value == val { 
		return curr
	} else {
		return nil
	}
}

func (l *List) Delete(val int) {
	curr := l.head
	var prev *Node = nil
	for curr.next != nil {
		if curr.value == val && l.head == curr {
			l.head = curr.next
			return
		} else if curr.value == val {
			prev.next = curr.next
			return
		} else {
			prev = curr
			curr = curr.next
		}
	}
	if curr.value == val { 
		prev.next = nil	
	}
}

func (l *List) Reverse() {
	curr := l.head; l.tail = l.head 
	var prev *Node = nil
	for curr.next != nil {
		//save next node before it is overwritten
		tempNext := curr.next 
		curr.next = prev 

		//make prev the current node before advancing to next node
		prev = curr		
		curr = tempNext
	}
	curr.next = prev
	l.head = curr
}

func (l *List) Print() {
	if l.head == l.tail {
		fmt.Println("[",l.head.value,"]")
		return
	}

	curr := l.head
	for curr.next != nil {
		fmt.Print(curr.value, "-->")
		curr = curr.next
	}
	fmt.Println(curr.value)
}

func randNode() *Node {
	rand.Seed(time.Now().UnixNano())
	return &Node{
		value: rand.Intn(100),
		next: nil,
	}
}

func newNode(val int) *Node {
	return &Node {
		value: val,
		next: nil,
	}
}

func main() {
	list := &List{
		head: nil,
		tail: nil,
	}

	var nodes = make([]*Node, 15)
	for i := range nodes {
		nodes[i] = newNode(i*10)
		list.Append(nodes[i])
	}

	list.Print()

	list.Reverse()
	list.Print()

	list.Delete(10)
	list.Delete(50)
	list.Delete(120)
	list.Delete(70)
	list.Print()
}