package task3

type Node struct {
	Value int
	Next  *Node
}

func removeElements(head *Node, val int) *Node {
	dummy := Node{Next: head}
	prev := &dummy
	for current := head; current != nil; current = current.Next {
		if current.Value == val {
			prev.Next = current.Next
		} else {
			prev = current
		}
	}
	return dummy.Next
}
