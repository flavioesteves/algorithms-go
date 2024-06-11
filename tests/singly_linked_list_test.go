package tests

import (
	"fmt"
	"testing"

	"github.com/flavioesteves/algorithms/sort"
)

func TestSinglyLinkedList(t *testing.T) {

	list := sort.NewSinglyLinkedList[int]()

	list.Append(5)
	list.Append(7)
	list.Append(9)

	expected, ok := list.Get(2)
	if !ok || expected != 9 {
		t.Errorf("Expected list.get(2) to be 9, got %v", expected)
	}

	removed, ok := list.RemoveAt(1)

	if !ok || removed != 7 {
		t.Errorf("Expected list.removeAt(1) to be 7, got %v", removed)
	}

	if list.Length() != 2 {
		t.Errorf("Expected list length to be 2, got %v", list.Length())
	}

	list.Append(11)

	removed, ok = list.RemoveAt(1)
	if !ok || removed != 9 {
		t.Errorf("Expected list.removeAt(1) to be 9, got %v", removed)
	}

	removed, ok = list.Remove(9)
	fmt.Println(ok)
	fmt.Println(removed)
	if ok {
		t.Errorf("Expected list.remove(9) to be nil (not found)")
	}

	removed, ok = list.RemoveAt(0)
	if !ok || removed != 5 {
		t.Errorf("Expected list.removeAt(0) to be 5, got %v", removed)
	}

	removed, ok = list.RemoveAt(0)
	if !ok || removed != 11 {
		t.Errorf("Expected list.removeAt(0) to be 11, got %v", removed)
	}

	if list.Length() != 0 {
		t.Errorf("Expected list length to be 0, got %v", list.Length())
	}

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	expected, ok = list.Get(2)
	if !ok || expected != 5 {
		t.Errorf("Expected list.get(2) to be 5, got %v", expected)
	}

	expected, ok = list.Get(0)
	if !ok || expected != 9 {
		t.Errorf("Expected list.get(0) to be 9, got %v", expected)
	}

	removed, ok = list.Remove(9)
	if !ok || removed != 9 {
		t.Errorf("Expected list.remove(9) to be 9")
	}

	if list.Length() != 2 {
		t.Errorf("Expected list length to be 2, got %v", list.Length())
	}

	expected, ok = list.Get(0)
	if !ok || expected != 7 {
		t.Errorf("Expected list.get(0) to be 7, got %v", expected)
	}

}
