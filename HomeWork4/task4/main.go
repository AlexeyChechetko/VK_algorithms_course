package task4

func maxMinMultiplication(data []int) int {
	if len(data) < 2 {
		return -1
	}
	minIndex, maxIndex := 0, 0
	for left := 2*minIndex + 1; left < len(data); left = 2*minIndex + 1 {
		minIndex = left
	}
	for right := 2*maxIndex + 2; right < len(data); right = 2*maxIndex + 2 {
		maxIndex = right
	}
	return data[minIndex] * data[maxIndex]
}
