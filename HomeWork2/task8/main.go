package task8

type Node struct {
	Value int
	Next  *Node
}

func mergeTwoLists(list1, list2 *Node) *Node {
	dummy := Node{}
	tail := &dummy
	for list1 != nil && list2 != nil {
		if list1.Value <= list2.Value {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}
	return dummy.Next
}
