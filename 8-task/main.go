// package main

// import "fmt"

// type Stack struct {
// 	items []int
// }

// func (s *Stack) Push(v int) {
// 	s.items = append(s.items, v)
// }

// func (s *Stack) Pop() (int, bool) {
// 	if len(s.items) < 0 {
// 		return 0, false
// 	}
// 	val := s.items[len(s.items)-1]
// 	s.items = s.items[:len(s.items)-1]
// 	return val, true

// }

// func (s *Stack) Peek() (int, bool) {
// 	if len(s.items) == 0 {
// 		return 0, false
// 	}
// 	return s.items[len(s.items)-1], true
// }

// func main() {

// 	stack := Stack{}
// 	stack.Push(10)
// 	stack.Push(20)
// 	stack.Push(30)
// 	fmt.Println("Stack after pushes:", stack.items)

// 	top, _ := stack.Peek()
// 	fmt.Println("Peek top:", top)

// 	popped, _ := stack.Pop()
// 	fmt.Println("Popped:", popped)
// 	fmt.Println("Stack after pop:", stack.items)
// }

package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack)Pop()(int , bool){
	if len(s.items)<0{
		return 0,false
	}
	val:=s.items[len(s.items)-1]
	s.items=s.items[:len(s.items)-1]
	return val ,true
}
func (s *Stack)Peek()(int , bool){
	if len(s.items)<0{
		return 0,false
	}
	val:=s.items[len(s.items)-1]
	return val ,true
}



func main() {
	s := Stack{}
	s.Push(20)
	s.Push(30)
	s.Push(50)
	s.Push(40)

	fmt.Println(s.items)
	i,_:=s.Pop()
	fmt.Println(i)
	v,_:=s.Peek()
	fmt.Println(v)
	

}
