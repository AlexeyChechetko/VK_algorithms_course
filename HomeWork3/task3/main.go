package task3

import "sort"

func feedAnimals(animals, food []int) int {
	if len(animals) == 0 || len(food) == 0 {
		return 0
	}
	sort.Ints(animals)
	sort.Ints(food)
	count := 0
	for _, portion := range food {
		if portion >= animals[count] {
			count++
		}
		if count == len(animals) {
			break
		}
	}
	return count
}
