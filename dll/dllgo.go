package main

import (
	"fmt"
	"os"
)
type Node struct{
	data int
	next *Node
	prev *Node

}
var head *Node
var tail *Node
func (ll *Node)Append(ele int){
	//var head *Node
	//var tail *Node
	createnode:=&Node{data:ele,next:nil,prev:nil}
	if head ==nil{
		head = createnode
		tail = createnode
	}else{
		createnode.prev=tail
		tail.next=createnode
		tail=createnode

	}


}
func (ll *Node)displayreverse(){
	current:=tail
	for current!=nil{
		fmt.Print(current.data," ")
		current = current.prev

	}
	fmt.Println()
}
func (ll *Node)displayforward(){
	current:=head
	for current!=nil{
		fmt.Print(current.data," ")
		current = current.next

	}
	fmt.Println()
}
func (ll *Node)Appendbegin(ele int){
	
	createnode:=&Node{data:ele,next:nil,prev:nil}
	if head ==nil{
		head = createnode
		tail = createnode
	}else{
		createnode.next=head
		head.prev=createnode
		head=createnode

	}


}


func main(){
	dll:=Node{}
	var choice ,ele int
	for true{
		fmt.Println("Enter your choice")
		fmt.Println("1.Insert node at end")
		fmt.Println("2.Insert node at beginning")
		fmt.Println("3.Traverse forward")
		fmt.Println("4.Traversw reverse")
		fmt.Println("5. exit")
		fmt.Scanln(&choice)
		switch choice {
		case 1 :fmt.Println("Enter element")
		        fmt.Scanln(&ele)
				dll.Append(ele) 
		case 2: fmt.Println("Enter element")
		        fmt.Scanln(&ele)
		        dll.Appendbegin(ele)
		case 3 : dll.displayforward()
		case 4: dll.displayreverse()
		case 5 : os.Exit(0)
		}
	}
}