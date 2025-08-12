package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) AddtoHead(val int) {
	NewNode := &Node{data: val, next: l.head}
	l.head = NewNode
}

func (l *LinkedList) AddTail(val int) {
	newNode := &Node{data: val}
	if l.head == nil {
		l.head = newNode
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

func (l *LinkedList) Delete(val int) {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	if l.head.data==val{
		l.head=l.head.next
		return
	}
	curr:=l.head
	if curr.next!=nil && curr.next.data!=val{
		curr=curr.next
	}
	if curr.next==nil{
		fmt.Println("Value not found")
		return
	}
	curr.next=curr.next.next
}

func (l *LinkedList) Display(){
	curr:=l.head
	for curr!=nil{
		fmt.Printf("%d <-",curr.data)
		curr=curr.next
	}
	fmt.Println("nil")
}

func main() {

	ll:=LinkedList{}
	ll.AddtoHead(20)
	ll.AddTail(40)
	ll.AddtoHead(10)
	ll.AddTail(30)
	ll.Display()
	ll.Delete(10)
	ll.Display()

}