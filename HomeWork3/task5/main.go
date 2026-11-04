package task5

func twoSum(data []int, target int) []int {
	cache := make(map[int]int)
	for i, value := range data {
		diff := target - value
		if index, found := cache[diff]; found {
			return []int{index, i}
		}
		// Добавляем после поиска, чтобы не использовать один индекс дважды.
		cache[value] = i
	}
	return []int{}
}
