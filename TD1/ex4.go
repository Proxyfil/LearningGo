package main

import (
	"fmt";
	"math/rand";
	"time";
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l LinkedList) Append(data int) LinkedList{
	if(l.head == nil){
		l.head = &Node{data: data}
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = &Node{data: data}
	}

	return l
}

func (l LinkedList) Print(){
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func (l LinkedList) Delete(data int) LinkedList{
	if l.head == nil {
		return l
	}
	if l.head.data == data {
		l.head = l.head.next
		return l
	}
	current := l.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return l
		}
		current = current.next
	}

	return l
}

func (l LinkedList) InsertAtPosition(data int, position int) LinkedList{
	if position == 0 {
		l.head = &Node{data: data, next: l.head}
		return l
	}
	current := l.head
	for i := 0; i < position-1; i++ {
		if current.next == nil {
			current.next = &Node{data: data}
			return l
		}
		current = current.next
	}
	if current == nil {
		return l
	}
	current.next = &Node{data: data, next: current.next}

	return l
}

func main() {
	rand.Seed(time.Now().UnixNano())

	list := LinkedList{}
	list = list.Append(1)
	list = list.Append(2)
	list = list.Append(3)
	list = list.Append(4)
	list.Print()

	list = list.Delete(2)
	list.Print()

	list = list.Delete(1)
	list.Print()

	list = list.InsertAtPosition(2, 5)
	list.Print()
}