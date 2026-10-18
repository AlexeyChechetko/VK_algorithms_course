package task1

func reversePartOfArray(arr []int, k int) {
	if len(arr) == 0 {
		return
	}
	k %= len(arr)
	if k == 0 {
		return
	}
	reverseArray(arr)
	reverseArray(arr[:k])
	reverseArray(arr[k:])
}

func reverseArray(arr []int) {
	for left, right := 0, len(arr)-1; left < right; left, right = left+1, right-1 {
		arr[left], arr[right] = arr[right], arr[left]
	}
}
