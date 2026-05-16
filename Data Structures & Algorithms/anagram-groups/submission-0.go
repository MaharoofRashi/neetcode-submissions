func groupAnagrams(strs []string) [][]string {

	// Make to byte, so easy to sort
	toSort := make([][]byte, len(strs))
	for i, s := range strs {
		toSort[i] = []byte(s)
	}

	// add it into a map after sorting 
	m := map[string][]string{}
	for i, s := range strs {
		sort.Slice(toSort[i], func (a, b int) bool {
			return toSort[i][a] < toSort[i][b]
		})

		key := string(toSort[i])
		m[key] = append(m[key], s)
	}

	// collect all map values and result in to result

	result := [][]string{}
	for _, group := range m {
		result = append(result, group)
	}

	return result
}
