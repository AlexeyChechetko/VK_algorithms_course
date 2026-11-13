package task2

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func isSymmetric(root *Node) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *Node) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Value == right.Value &&
		isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}
