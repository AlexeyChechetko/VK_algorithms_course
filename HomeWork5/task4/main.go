package task4

// Структура узла дерева
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Функция поиска kого наименьшего элемента
func kthSmallest(root *Node, k int) (int, bool) {
	return kthElement(root, k, false)
}

// Функция поиска kого наибольшего элемента
func kthLargest(root *Node, k int) (int, bool) {
	return kthElement(root, k, true)
}

// Основная функция
func kthElement(root *Node, k int, largest bool) (int, bool) {
	// Если k < 1, то данные введены неверно
	if k < 1 {
		return 0, false
	}
	// Создаем стек узлов дерева, представленный динамическим массивом
	stack := make([]*Node, 0)
	// Устанавливаем текущий обрабатываем узел в корень дерева
	current := root
	// Основной цикл, выполняющийся пока мы обрабатываем вершины или стек не пуст,
	// то есть еще есть необработанные вершины
	for current != nil || len(stack) > 0 {
		// Каждый раз спускаемся в самое лево (при поиске наименьшего элемента) или право (наибольшего)
		// Добавляем элементы в стек
		for current != nil {
			stack = append(stack, current)
			if largest {
				current = current.Right
			} else {
				current = current.Left
			}
		}
		// Достаем элемент с вершины стека, с ним будем работать
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// Вычитаем из k единицу, так как это какой-то по счету обрабатываемый элемент
		k--
		// Если k = 0, то мы нашли kый наименьший или наибольший элемент
		if k == 0 {
			return current.Value, true
		}
		// Если мы пока не нашли, то мы должны проверить у текущего "самого правого" узла нет ли левого поддерева
		// в случае поиска наибольшего элемента. Если есть, то мы заново обрабатываем его по циклу, спускаясь вниз вправо и т.д.
		// Для поиска наименьшего элемента аналогично
		if largest {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return 0, false
}
