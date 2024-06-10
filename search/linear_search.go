package search

func LinearSearch(haystack []int, needle int) bool {

	for _, element := range haystack {
		if element == needle {
			return true
		}
	}
	return false
}
