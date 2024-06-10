package tests

import (
	"github.com/flavioesteves/algorithms/search"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	tests := []struct {
		name   string
		target int
		want   bool
	}{
		{name: "Find existing element (69)", target: 69, want: true},
		{name: "Find existing element (69420)", target: 69420, want: true},
		{name: "Find existing element (1)", target: 1, want: true},
		{name: "Not Found (non-existent elementr)", target: 69421, want: false},
		{name: "Not Found (1336)", target: 1336, want: false},
		{name: "Not Found (element less that first)", target: 0, want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := search.LinearSearch(foo, tc.target)
			if got != tc.want {
				t.Errorf("Test case: %s - expected %v, got %v", tc.name, tc.want, got)
			}
		})
	}
}
