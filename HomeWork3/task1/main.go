package task1

func binarySearchSqrt(target int) int {
	if target < 0 {
		return -1
	}
	if target < 2 {
		return target
	}
	left, right := 1, target/2+1
	for left <= right {
		middle := (left + right) / 2
		if middle*middle <= target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return right
}
