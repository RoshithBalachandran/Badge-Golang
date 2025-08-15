package main

import "fmt"

type Queue struct {
	item []int
}

func (q *Queue) Enqueue(val int) {
	q.item = append(q.item, val)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.item) == 0 {
		fmt.Println("List is empty")
		return 0,false
	}
	val:=q.item[0]
	q.item=q.item[1:]
	return val,true
}

func (q *Queue)Peek(){
	if len(q.item) == 0 {
		fmt.Println("List is empty")
		return 
	}
	fmt.Println("Peek :",q.item[0])
}

func (q *Queue)Display(){
	if len(q.item)==0{
		fmt.Println("List is empty")
	}
	fmt.Println("Queue :",q.item)
}

func main() {
	var q Queue
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Display()
	q.Peek()
	q.Display()
	q.Dequeue()
	q.Display()

}