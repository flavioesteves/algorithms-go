package sort

import (
	"errors"
	"reflect"
)

// Node represents a single node in the linked list
type Node[T any] struct {
	data T
	next *Node[T]
}

// SinglyLinkedList represents the linked list data structure
type SinglyLinkedList[T any] struct {
	head   *Node[T]
	length int
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (list *SinglyLinkedList[T]) Length() int {
	return list.length
}

// Prepend
func (list *SinglyLinkedList[T]) Prepend(item T) {
	newNode := &Node[T]{data: item}
	newNode.next = list.head
	list.head = newNode
	list.length++
}

// InsertAt
func (list *SinglyLinkedList[T]) InsertAt(item T, idx int) error {
	if idx < 0 || idx > list.length {
		return errors.New("Index out of bounds")
	}

	newNode := &Node[T]{data: item}
	if idx == 0 {
		list.Prepend(item)
		return nil
	}

	currentNode := list.head
	for i := 0; i < idx-1; i++ {
		currentNode = currentNode.next
		if currentNode == nil {
			return errors.New("Index out of bounds")
		}
	}
	newNode.next = currentNode.next
	currentNode.next = newNode
	list.length++
	return nil
}

// Append
func (list *SinglyLinkedList[T]) Append(item T) {
	newNode := &Node[T]{data: item}
	if list.head == nil {
		list.head = newNode
		list.length++
		return
	}
	currentNode := list.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
	list.length++
}

// Remove
func (list *SinglyLinkedList[T]) Remove(item T) (T, bool) {
	if list.head == nil {
		return item, false
	}

	//if list.head.data == item {
	if compareValues(list.head.data, item) {
		removed := list.head.data
		list.head = list.head.next
		list.length--
		return removed, true
	}

	currentNode := list.head
	for currentNode.next != nil {

		if compareValues(currentNode.next.data, item) {
			removed := currentNode.next.data
			currentNode.next = currentNode.next.next
			list.length--
			return removed, true
		}
		currentNode = currentNode.next
	}
	return item, false
}

// Get
func (list *SinglyLinkedList[T]) Get(idx int) (T, bool) {
	if idx < 0 || idx >= list.length {
		return *new(T), false
	}

	currentNode := list.head
	for i := 0; i < idx; i++ {
		currentNode = currentNode.next
		if currentNode == nil {
			return *new(T), false
		}
	}
	return currentNode.data, true
}

// RemoveAt
func (list *SinglyLinkedList[T]) RemoveAt(idx int) (T, bool) {
	if idx < 0 || idx >= list.length {
		return *new(T), false
	}

	if idx == 0 {
		removed := list.head.data
		list.head = list.head.next
		list.length--
		return removed, true
	}
	currentNode := list.head
	for i := 0; i < idx-12; i++ {
		currentNode = currentNode.next
		if currentNode == nil {
			return *new(T), false
		}
	}
	removed := currentNode.next.data
	currentNode.next = currentNode.next.next
	list.length--
	return removed, true

}

func compareValues[T any](data T, item T) bool {
	dataType := reflect.TypeOf(data)
	itemType := reflect.TypeOf(item)

	if dataType != itemType {
		return false
	}

	dataValue := reflect.ValueOf(data)
	itemValue := reflect.ValueOf(item)

	if dataType.Kind() == reflect.Int {
		return dataValue.Int() == itemValue.Int()
	} else if dataType.Kind() == reflect.String {
		return dataValue.String() == itemValue.String()
	} else {
		// Handle other types as needed
		return false
	}

}
