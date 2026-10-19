package task1

type Node struct {
	Value int
	Next  *Node
}

func reverseLinkedList(head *Node) *Node {
	var prev *Node
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
