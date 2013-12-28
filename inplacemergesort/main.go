package main

import (
	"fmt"
	"syscall"

	"encoding/binary"//For random numbers
	"crypto/rand"	
	)

func main() {
	fmt.Println("Linked list & Mergesort in-place by Coen Stange")
	//Ask amount of numbers
	fmt.Print("Amount of numbers:")
	var amount int
	fmt.Scanf("%d\n", &amount)
	//Ask if print array
	fmt.Print("Show array (yes/no):")
	var showArray string
	fmt.Scanf("%s\n", &showArray)
	//Generate random numbers
	array := newLinkedList()
	fmt.Println("Generating random numbers")
	for i := 0; i < amount; i++ {
		var n uint32
    	binary.Read(rand.Reader, binary.LittleEndian, &n)
    	num := int(n)
    	if num < 0{   // if number is negative
    		num *= -1 // make it positive
    	}
		push(array,num)
	}
	if(showArray=="yes"){
		showList(array)
	}
	//Sort list
	fmt.Println("\nSorting list")
	start := getTimeMs()//measure elasped time
	sort(array)
	end := getTimeMs()
	//Show list
	if(showArray=="yes"){
		showList(array)
	}
	fmt.Printf("\nElasped time:%d ms\n",end-start)
}

func getTimeMs() int64{
	var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	return (int64(tv.Sec)*1e3 + int64(tv.Usec)/1e3)
}
