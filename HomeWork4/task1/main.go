package task1

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func buildTree(data []int) *Node {
	return buildTreeFromIndex(data, 0)
}

func buildTreeFromIndex(data []int, index int) *Node {
	if index >= len(data) {
		return nil
	}
	return &Node{
		Value: data[index],
		Left:  buildTreeFromIndex(data, 2*index+1),
		Right: buildTreeFromIndex(data, 2*index+2),
	}
}
