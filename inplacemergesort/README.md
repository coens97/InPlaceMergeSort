In-place merge sort using linked list with Go
=========
Less memory usage by not creating a temporary array
How does it work?
---
The list is split down in 3 parts:: sorted, left, right.
![](http://i.imgur.com/13I6gNw.png) 

If the first value of the left list is smaller then the beginning of the left list is moved one to the right. 
![](http://i.imgur.com/2mGCSfz.png) 

If the first value of the right list is smaller then change the pointer from the last node to the second node from the second list. Change the pointer from the last node from the sorted to the node which was the first of the right list and change the pointer of that to the first of the left list. 

![](http://i.imgur.com/AYOAiHz.png) 

When the left or right list ks empty the part of the list is sorted
How to run?
---
 - Install The Go Programming Language
 - ```go run main.go linkedlist.go mergesort.go```
