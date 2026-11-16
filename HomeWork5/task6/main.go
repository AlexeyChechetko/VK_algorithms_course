package task6

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func mirrorTree(root *Node) *Node {
	if root == nil {
		return nil
	}
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}
