package tests

import (
	"github.com/flavioesteves/algorithms/sort"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	expected := []int{3, 4, 7, 9, 42, 69, 420}

	sort.BubbleSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected sorted array to be %v, got %v", expected, arr)
	}

}
