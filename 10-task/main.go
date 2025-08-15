package main

import "fmt"

 func Sent (ch chan string)  {
		fmt.Println("Sending channel")
		ch<-"Hellow World"
		fmt.Println("Message Sent........")
	}

func main() {
	ch := make(chan string)
	go Sent(ch)
	fmt.Println("Waiting for channels")
	msg:=<-ch
	fmt.Println("MEssage :",msg)

}