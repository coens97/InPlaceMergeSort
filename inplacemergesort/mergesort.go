package main 

func sort(linked *Node) {
	mergesort(linked,linked.value)
}

func mergesort(start *Node,length int) *Node{
	if length > 1{
		var leftLength, rightLength int
		leftStart := start
		// Split the lists
		leftLength = length / 2
		rightLength = length - leftLength
		// sort both sides
		rightStart := mergesort(leftStart, leftLength)		
		mergesort(rightStart, rightLength)
		// merge
		for  0!=leftLength&&0!=rightLength{
			if leftStart.next.value < rightStart.next.value{
				leftStart = leftStart.next
				leftLength--
			} else{
				moveNode(rightStart,leftStart)
				leftStart = leftStart.next
				rightLength--
			}
		}
		//return end of linkedlist
		leng := leftLength + rightLength
		for leng>0{
			leftStart = leftStart.next
			leng--
		}
		return leftStart
	}
	return start.next
}