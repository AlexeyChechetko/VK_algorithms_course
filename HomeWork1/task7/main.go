package task7

func moveZeroes(nums []int) {
	next := 0
	for current := range nums {
		if nums[current] != 0 {
			nums[next], nums[current] = nums[current], nums[next]
			next++
		}
	}
}
