package search

import (
	"math"
)

func BsList(haystack []int, needle int) bool {
	low := float64(0)
	high := float64(len(haystack))

	for low < high {
		midPoint := int(math.Floor(low + (high-low)/2))
		value := int(haystack[int(midPoint)])

		if value == needle {
			return true
		} else if value > needle {
			high = float64(midPoint)
		} else {
			low = float64(midPoint) + 1
		}
	}
	return false
}
