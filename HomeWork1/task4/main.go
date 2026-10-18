package task4

func sortBinaryArray(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] == 0 {
			left++
		} else if arr[right] == 1 {
			right--
		} else {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
}
