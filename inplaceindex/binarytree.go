package main

import (
        "fmt"
        )
type bNode struct{
        value *Node//points to node in linked list
        left *bNode
        right *bNode
}

type Queue struct{
        begin *QueueNode
        end *QueueNode
}

type QueueNode struct{
        value *bNode
        next *QueueNode
}

func newBNode() *bNode{
        return &bNode{value:nil,left:nil,right:nil}
}

func indexList(start *Node) *bNode{
        //create binary tree
        bTree := newBNode()
        //insert binary tree
        levelOrderInsert(bTree, start.next, 1)
        return bTree
}

func levelOrderInsert(root *bNode, list *Node, level int){
        //check if non nil values are given
        if(list==nil||root==nil){
                return
        }
        //set the value
        root.value = list
        //loop trough list
        for i := 0; i < level; i++ {
                list = list.next
                if(list == nil){
                        return
                }
        }
        //insert left node
        root.left = newBNode()
        levelOrderInsert(root.left,list,level*2)
        //insert right node
        for i := 0; i < level; i++ {
                list = list.next
                if(list == nil){
                        return
                }
        }
        root.right = newBNode()
        levelOrderInsert(root.right,list,level*2)
}

func newQueue(val *bNode) *Queue{
        node := newQueueNode(val)
        return &Queue{begin:node,end:node}
}

func newQueueNode(val *bNode) *QueueNode{
        return &QueueNode{value:val,next:nil}
}

func printTree(root *bNode){
        queue := newQueue(root)

        i := 0
        for queue.begin!=nil{
                //show value
                if(i % 4 == 0){//if first column
                        fmt.Print("|")
                }
                fmt.Printf("%5d || %10d",i,queue.begin.value.value.value)
                if(i % 4 == 3){ //after the last colum do new line
                        fmt.Print("|\n")
                }
                //Add left and right node to queue
                if(queue.begin.value.left!=nil){
                        queue.end.next = newQueueNode(queue.begin.value.left)
                        queue.end = queue.end.next
                }
                if(queue.begin.value.right!=nil){
                        queue.end.next = newQueueNode(queue.begin.value.right)
                        queue.end = queue.end.next
                }
                queue.begin = queue.begin.next
                i++
        }
}
