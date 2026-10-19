package task2

type Node struct {
	Value int
	Next  *Node
}

func middleNode(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
