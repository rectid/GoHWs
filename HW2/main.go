package main

import (
	"awesomeProject/collection/LinkedList"
	"fmt"
)

func main() {
	list := LinkedList.New(4)
	list.UpdateAt(0, 123)
	fmt.Println(list.Size())
	list.Add(2)
	list.UpdateAt(1, 321)
	fmt.Println(list.At(4))
	fmt.Println(list.At(0))
	list.Pop()
	list.DeleteFrom(1)
	fmt.Println(list.Size())
	list = LinkedList.NewFromSlice([]int{1, 2, 3, 4})
	fmt.Println(list.Size())
	fmt.Println(list.At(2))
	list.DeleteFrom(1)
	fmt.Println(list.At(1))

}
