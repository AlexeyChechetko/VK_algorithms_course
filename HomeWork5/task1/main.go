package task1

// Структура узла дерева
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Функция проверки полноты дерева по определению 1
func isComplete(root *Node) bool {
	// Если дерево пустое, то считаем его полным
	if root == nil {
		return true
	}
	// Будем выполнять BFS в бинарном дереве, проходясь по его уровням и фиксируя пропуски узлов
	// level - список узлов на текущем уровне. На первом уровне у нас только один узел - корень
	level := []*Node{root}
	// Переменная gap принимает значение true, когда встречается пропуск узла в дереве
	gap := false
	// Пока у нас есть узлы выполняем цикл
	for len(level) > 0 {
		// Массив узлов следующего уровня
		nextLevel := make([]*Node, 0)
		for _, node := range level {
			// Для каждого узла текущего уровня смотрим его потомков и проверям пустые они или нет
			for _, child := range [2]*Node{node.Left, node.Right} {
				// Если потомок пустой, то фиксируется пропуск, gap = true
				// Если при обработке очередного узла после фиксации пропуска встретиться узел, то это уже не полное бинарное дерево
				if child == nil {
					gap = true
					continue
				}
				if gap {
					return false
				}
				nextLevel = append(nextLevel, child)
			}
		}
		level = nextLevel
	}
	return true
}
