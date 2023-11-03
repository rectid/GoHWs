package LinkedList

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
}

func New(size int) *LinkedList {
	if size <= 0 {
		panic("Size must be greater than 0")
	}
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

func (l *LinkedList) At(pos int) (int, error) {
	if pos < 0 {
		panic("Index out of range")
	}
	currentNode := l.head
	for i := 0; i < pos; i++ {
		if currentNode.next == nil {
			panic("Index out of range")
		}
		currentNode = currentNode.next
	}
	return currentNode.val, nil
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
	if pos < 0 {
		panic("Index out of range")
	}
	currentNode := l.head
	for i := 0; i < pos; i++ {
		if currentNode.next == nil {
			panic("Index out of range")
		}
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
	if pos < 0 {
		panic("Index out of range")
	}
	currentNode := l.head
	for i := 0; i < pos-1; i++ {
		if currentNode.next == nil {
			panic("Index out of range")
		}
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
}

func NewFromSlice(s []int) *LinkedList {
	if len(s) == 0 {
		panic("Size of slice must be greater than 0")
	}
	head := node{val: s[0], next: nil}
	currentNode := &head
	for i := 1; i < len(s); i++ {
		currentNode.next = &node{val: s[i], next: nil}
		currentNode = currentNode.next
	}

	return &LinkedList{head: &head}
}
