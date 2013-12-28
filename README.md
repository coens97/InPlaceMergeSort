In-place merge sort faster then built-in sort go
=========
Is faster but using more memory
---
 * Go built-in array (or slice) use less memory then an linked list
 * Go sort implementation uses more sorting algorithms at the same time and is not always recursive 

In-place merge sort is using less memory then the regular merge sort because that makes temporary arrays
 
Speed results
---
The time in milliseconds to finish sorting for a given number of elements

Number of elements | Sort (ms) | In-place merge sort (ms)
---|---|---
1E3 | 1 | 0
5E3 | 3 | 1
1E4 | 6 | 2
5E4 | 31 | 9
1E5 | 70 | 20
5E5 | 414 | 156
1E6 | 901 | 393
5E6 | 5241 | 3264
1E7 | 10800 | 7879
5E7 | 60791 |  56471