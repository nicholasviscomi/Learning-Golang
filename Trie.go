package main

import (
	"fmt"
)

const (
	numChars = 256 //number of possible ASCII values
)

// https://www.javatpoint.com/trie-data-structure

type tNode struct {
	val rune //char
	children []*tNode
	terminal bool //signifies this is the last letter in a word
}

func newNode(val rune, terminal bool) *tNode {
	return &tNode {
		//length of zero but 256 capacity
		val: val,
		children: make([]*tNode, numChars),
		terminal: terminal,
	}
}

func (head *tNode) insert(str string) {
	/*
	-go through each character in str
	---if a node with the letter already exists enter that node
	---if not, create a node with that letter and enter it
	*/
	curr := head
	for _, char := range str {
		n := byte(char); terminal := false
		//put a new node at the index of the current char
		//OPTIMIZE HERE: makes extra function calls if the value already exists
		if curr.children[n] == nil {
			curr.children[n] = newNode(char, terminal) 
		}
		curr = curr.children[n]
	}
	//set the last letter's node as the end of the word
	curr.terminal = true
}

func (head *tNode) search(str string) bool {
	/*
	-go through each character in str
	---if a node with the letter already exists enter that node
	---if not, return false
	-if you get through the whole str without returning, the string exists
	*/
	curr := head
	for _, char := range str {
		n := byte(char)
		if curr.children[n].val == char {
			curr = curr.children[n]	
		} else {
			return false
		}
	}
	//if the last letter's node is the end of a word
	if curr.terminal {
		return true
	} else {
		return false
	}

}

func (node *tNode) hasChildren() bool {
	return len(node.children) > 0
}

func (head *tNode) delete(str string) {
	traversedNodes := make([]*tNode, len(str))

	curr := head
	for _, char := range str {
		traversedNodes = append(traversedNodes, curr)
		n := byte(char)
		if curr.children[n].val == char {
			curr = curr.children[n]	
		}
	}
	if !curr.terminal { return }
	if curr.hasChildren() {
		//there are more words after this final letter
		//don't delete all the traversed nodes
		curr.terminal = false
	} else {
		//no other words after this
		//delete all the traversed nodes
		for _, node := range traversedNodes {
			node.children = nil
		}
	}
}

func main() {
	head := newNode('0', false)
	head.insert("hello")
	head.insert("apple")
	head.insert("app")
	head.insert("snore")
	head.insert("snores")
	fmt.Println("hello =", head.search("hello"))
	fmt.Println("apple =", head.search("apple"))
	fmt.Println("app =", head.search("app"))
	fmt.Println("snore =", head.search("snore"))
	fmt.Println("snores =", head.search("snores"))
	head.delete("snore")
	head.delete("snores")
	fmt.Println("snore =", head.search("snore"))
	fmt.Println("snores =", head.search("snores"))
}