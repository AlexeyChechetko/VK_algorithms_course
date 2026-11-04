package task2

func copyTime(n, x, y int64) int64 {
	if n <= 0 {
		return 0
	}
	if x <= 0 || y <= 0 {
		return -1
	}
	first := min(x, y)
	remaining := n - 1
	left, right := int64(0), remaining*first
	for left < right {
		middle := (left + right) / 2
		if middle/x+middle/y >= remaining {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return first + left
}
