package task6

func evenFirst(arr []int) {
	evenIndex := 0
	for current := range arr {
		if arr[current]%2 == 0 {
			arr[evenIndex], arr[current] = arr[current], arr[evenIndex]
			evenIndex++
		}
	}
}
