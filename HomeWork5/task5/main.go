package task5

type Node struct {
	Value         int
	Left          *Node
	Right         *Node
	BalanceFactor int
}

func fillBalanceFactors(root *Node) {
	fillHeight(root)
}

func fillHeight(node *Node) int {
	if node == nil {
		return -1
	}
	leftHeight := fillHeight(node.Left)
	rightHeight := fillHeight(node.Right)
	node.BalanceFactor = leftHeight - rightHeight
	return 1 + max(leftHeight, rightHeight)
}
