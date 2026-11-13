package task5

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func isSameTree(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Value == b.Value &&
		isSameTree(a.Left, b.Left) &&
		isSameTree(a.Right, b.Right)
}
