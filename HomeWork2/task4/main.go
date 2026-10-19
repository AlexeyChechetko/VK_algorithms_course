package task4

func isSubsequence(a, b string) bool {
	i := 0
	for j := 0; i < len(a) && j < len(b); {
		if a[i] == b[j] {
			i += 1
		}
		j += 1
	}
	return i == len(a)
}
