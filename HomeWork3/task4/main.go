package task4

func extraLetter(a, b string) string {
	counts := make(map[byte]int)
	for i := 0; i < len(a); i++ {
		counts[a[i]]++
	}
	for i := 0; i < len(b); i++ {
		letter := b[i]
		if counts[letter] == 0 {
			return string(letter)
		}
		counts[letter]--
	}
	return ""
}
