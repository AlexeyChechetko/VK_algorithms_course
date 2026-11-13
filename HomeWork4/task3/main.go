package task3

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func minDepth(root *Node) int {
	if root == nil {
		return 0
	}
	level := []*Node{root}
	depth := 1
	for len(level) > 0 {
		nextLevel := make([]*Node, 0)
		for _, node := range level {
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		level = nextLevel
		depth++
	}
	return 0
}
