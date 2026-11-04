package task6

func groupAnagrams(strs []string) [][]string {
	groups := make([][]string, 0)
	indices := make(map[[26]int]int)
	for _, word := range strs {
		var key [26]int
		for i := 0; i < len(word); i++ {
			key[word[i]-'a']++
		}
		index, found := indices[key]
		if !found {
			index = len(groups)
			indices[key] = index
			groups = append(groups, nil)
		}
		groups[index] = append(groups[index], word)
	}
	return groups
}
