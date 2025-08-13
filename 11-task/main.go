package main


import "fmt"

func SendChannel(ch chan string){
	fmt.Println("Sending Channels")
	ch<-"Roshith"
	fmt.Println("Send Channels")
}


func main(){
	ch:=make(chan string)
	go SendChannel(ch)
	msg:=<-ch
	fmt.Println("Responce :",msg)

}