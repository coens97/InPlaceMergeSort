package main

import (
	"fmt"
	)

type Node struct {
	value int
	next *Node
}

func newLinkedList() *Node{
	return &Node{value:0,next:nil} //In first node store the length
}

func push(linked *Node,val int) {
	tmpNode := linked.next
	linked.next = &Node{value:val,next:tmpNode}
	linked.value++;
}

func moveNode(from *Node,to *Node){
	//change the pointers
	tmpNode := from.next //the node which will be moved
	//slice the node
	from.next = from.next.next
	//add the node
	tmpNext := to.next
	to.next = tmpNode
	tmpNode.next = tmpNext
}

func showList(linked *Node) {
	node := linked.next
	i := 0
	for node != nil{//go through list
		if(i % 4 == 0){//if first column
			fmt.Print("|")
		}
		fmt.Printf("%5d || %10d",i,node.value)
		if(i % 4 == 3){ //after the last colum do new line
			fmt.Print("|\n")
		}
		node = node.next//get next node of linkedlist
		i++
	}
}