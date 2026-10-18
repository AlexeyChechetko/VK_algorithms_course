package task3

func merge(arr1 []int, arr2 []int) []int {
	i, j := len(arr1)-len(arr2)-1, len(arr2)-1
	for dst := len(arr1) - 1; j >= 0; dst-- {
		if i >= 0 && arr1[i] > arr2[j] {
			arr1[dst] = arr1[i]
			i--
		} else {
			arr1[dst] = arr2[j]
			j--
		}
	}
	return arr1
}
