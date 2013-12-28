package main 

func sort(linked *Node) {
	mergesort(linked,linked.value)
}

func mergesort(start *Node,length int){
	if length > 1{
		var leftLength, rightLength int
		leftStart := start
		// Split the lists
		leftLength = length / 2
		rightLength = length - leftLength
		// sort both sides
		mergesort(leftStart, leftLength)
		// get beginning of right list
		rightStart := leftStart
		for i := 0; i < leftLength; i++ {
			rightStart = rightStart.next
		}
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
	}
}