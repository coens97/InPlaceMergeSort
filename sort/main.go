package main

import (
	"fmt"
	"sort"
	"syscall"

	"encoding/binary"//For random numbers
	"crypto/rand"	
)

func main() {
	fmt.Println("Builtin sorting")
	//Ask amount of numbers
	fmt.Print("Amount of numbers:")
	var amount int
    fmt.Scanf("%d\n", &amount)
    //Ask if print array
	fmt.Print("Show array (yes/no):")
	var showArray string
	fmt.Scanf("%s\n", &showArray)
    //Generate random numbers
    var array []int
    array = make([]int, amount, amount)
	fmt.Println("Generating random numbers")
	for i := 0; i < amount; i++ {
		var n uint32
    	binary.Read(rand.Reader, binary.LittleEndian, &n)
    	num := int(n)
    	if num < 0{   // if number is negative
    		num *= -1 // make it positive
    	}
		array[i] = num
	}
	if(showArray=="yes"){
		fmt.Println(array)
	}
	start := getTimeMs()//measure elasped time
	fmt.Println("\nSorting list")
	sort.Ints(array)
	end := getTimeMs()
	if(showArray=="yes"){
		fmt.Println(array)
	}
	fmt.Printf("\nElasped time:%d ms\n",end-start)
}

func getTimeMs() int64{
	var tv syscall.Timeval
	syscall.Gettimeofday(&tv)
	return (int64(tv.Sec)*1e3 + int64(tv.Usec)/1e3)
}
