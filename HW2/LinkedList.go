package main

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
}

func New(size int) *LinkedList {
	head := node{}
	currentNode := &head
	for i := 1; i < size; i++ {
		currentNode.next = &node{}
		currentNode = currentNode.next
	}

	return &LinkedList{head: &head}
}

func (l *LinkedList) Size() int {
	size := 0
	currentNode := l.head
	for currentNode != nil {
		size++
		currentNode = currentNode.next
	}
	return size
}

func (l *LinkedList) At(pos int) int {
	currentNode := l.head
	for i := 0; i < pos; i++ {
		currentNode = currentNode.next
	}
	return currentNode.val
}

func (l *LinkedList) Add(val int) {
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = &node{
		val:  val,
		next: nil,
	}
}

func (l *LinkedList) UpdateAt(pos int, val int) {
	currentNode := l.head
	for i := 0; i < pos; i++ {
		currentNode = currentNode.next
	}
	currentNode.val = val
}

func (l *LinkedList) Pop() {
	currentNode := l.head
	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = nil
}

func (l *LinkedList) DeleteFrom(pos int) {
	currentNode := l.head
	for i := 0; i < pos-1; i++ {
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
}

func NewFromSlice(s []int) *LinkedList {
	head := node{val: s[0], next: nil}
	currentNode := &head
	for i := 1; i < len(s); i++ {
		currentNode.next = &node{val: s[i], next: nil}
		currentNode = currentNode.next
	}

	return &LinkedList{head: &head}
}
