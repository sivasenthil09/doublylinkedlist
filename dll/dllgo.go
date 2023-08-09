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
var Head *Node
var Tail *Node
func (ll *Node)Append(ele int){
	//var head *Node
	//var tail *Node
	createnode:=&Node{data:ele,next:nil,prev:nil}
	if Head ==nil{
		Head = createnode
		Tail = createnode
	}else{
		createnode.prev=Tail
		Tail.next=createnode
		Tail=createnode

	}


}
func (ll *Node)displayreverse(){
	current:=Tail
	for current!=nil{
		fmt.Print(current.data," ")
		current = current.prev

	}
	fmt.Println()
}
func (ll *Node)displayforward(){
	current:=Head
	for current!=nil{
		//current = current.next
		fmt.Print(current.data," ")
		current = current.next

	}
	fmt.Println()
}
func (ll *Node)Appendbegin(ele int){
	
	createnode:=&Node{data:ele,next:nil,prev:nil}
	if Head ==nil{
		Head = createnode
		Tail = createnode
	}else{
		createnode.next=Head
		Head.prev=createnode
		Head=createnode

	}


}
func removeele(head *Node,ele int){
         for head!=nil{ 
			if head.data==ele && head.prev==nil{
              Head = head.next
			  return
			}
			if head.data==ele &&head.next==nil{
				newl:=head.prev
				newl.next=nil
				return
			}
			if head.data == ele{

			head=head.next
			newn:=head
			head=head.prev
			head=head.prev
			head.next=newn
			newn.prev=head

		  }
		  head = head.next
		}
}
func removebegin(head*Node){
	Head=head.next
}
func removeend(head*Node){
	for head!=nil{
		if head.next==nil{
			news:=head.prev
			news.next=nil
			return
		}
		head=head.next
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
		fmt.Println("8. exit")
		fmt.Println("5. remove element")
		fmt.Println("6. remove element at begining")
		fmt.Println("7. remove element at end")
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
		case 8 : os.Exit(0)
		case 5: fmt.Println("enter a element to remove")
		        fmt.Scanln(&ele)
				removeele(Head,ele)
		case 6: removebegin(Head)
		case 7: removeend(Head)
		}
	}
}