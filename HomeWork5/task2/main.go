package task2

import "container/heap"

// Документация: https://pkg.go.dev/container/heap#section-documentation

// Создаем структуру с 3 элеменатами.
// value - основное значение
// array - номер массива, из которого значение взято
// index - индекс этого значения в массиве array
type item struct {
	value int
	array int
	index int
}

// Создаем тип - массив  наших структур
type minHeap []item

// Данный тип должен реализовывать интерфейс из пакета container/heap.
// Реализуем необходимые функции в соответствие с их определением
func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(value any)    { *h = append(*h, value.(item)) }
func (h *minHeap) Pop() any {
	last := len(*h) - 1
	value := (*h)[last]
	*h = (*h)[:last]
	return value
}

// Основная функция, реализующая алгоритм
func mergeSortedArrays(arrays [][]int) []int {
	// Создаем массив наших структур
	h := make(minHeap, 0)
	// Заполняем его первыми элементами из наших массивов (наименьшими), их k штук
	for i, array := range arrays {
		if len(array) > 0 {
			h = append(h, item{value: array[0], array: i, index: 0})
		}
	}
	// Создаем неубывающую пирамиду их нашего созданного массива
	heap.Init(&h)
	// Создаем результирующий массив, куда будут записаны наши отсортированные значения
	result := make([]int, 0)
	// Пока в пирамиде есть элементы цикл выполняется
	for h.Len() > 0 {
		// Достаем минимальный элемент из пирамиды
		current := heap.Pop(&h).(item)
		// Добавляем его в наш результирующий массив
		result = append(result, current.value)
		// Из массива этого элемента будет браться следующий
		next := current.index + 1
		// Если в массиве еще остались необработанные элементы то добавляем следующий в пирамиду (таким образом в ней всегда будет k элементов)
		if next < len(arrays[current.array]) {
			heap.Push(&h, item{
				value: arrays[current.array][next],
				array: current.array,
				index: next,
			})
		}
	}
	return result
}
