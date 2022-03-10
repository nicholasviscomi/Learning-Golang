package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Working resizable stack with integer values
*/

type Stack struct {
	values []int
	capacity int
	top int //index for the top value
}

func (s *Stack) init(size int) {
	s.capacity = size
	s.values = make([]int, size)
	s.top = -1
}

func (s *Stack) resize(size int) {
	if !s.isFull() || size <= 0 { return }

	if size < s.capacity {
		s.values = s.values[0:size]
		s.top = size - 1
	} else if size > s.capacity {
		newVals := make([]int, size)
		copy(newVals, s.values)
		s.values = newVals
	}
	s.capacity = size
}

func (s *Stack) push(val int) {
	if s.isFull() {
		fmt.Println("Stack is already full")
		return
	}
	s.top = s.top + 1
	s.values[s.top] = val
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		return s.top //-1
	}
	res := s.values[s.top]
	s.values[s.top] = 0
	s.top--
	
	return res
}

func (s *Stack) peek() int {
	return s.values[s.top]
}

func (s *Stack) isFull() bool {
	return (s.top + 1) == s.capacity
}

func (s *Stack) isEmpty() bool {
	return (s.top == -1)
}

func (s *Stack) size() int {
	return s.top + 1
}

func (s *Stack) available() int {
	return s.capacity - s.size()
}

func (s *Stack) print() {
	fmt.Println(s.values)
}

func main() {
	fmt.Println("-----Stack-----")
	rand.Seed(time.Now().UnixNano())
	var s Stack;
	s.init(20)
	for !s.isFull() {
		s.push(rand.Intn(100))
	}
	s.print()

	s.resize(10)
	s.print()
	s.pop()
	s.print()
	s.push(rand.Intn(100))
	s.print()

	s.resize(20)
	s.print()
	s.push(rand.Intn(100))
	s.print()
	s.pop()
	s.print()
}