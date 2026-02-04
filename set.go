package collections

import (
	"slices"
)

type Set[T comparable] map[T]struct{}

// NewSet creates a new set and adds all the items
func NewSet[T comparable](items ...T) Set[T] {
	s := Set[T]{}
	s.Add(items...)
	return s
}

// Add adds the items to the set
func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

// AddAll adds all items from the other set
func (s Set[T]) AddAll(items Set[T]) {
	for item := range items {
		s[item] = struct{}{}
	}
}

// Remove removes the items from the set
func (s Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

// RemoveAll removes all items from the other set
func (s Set[T]) RemoveAll(items Set[T]) {
	for item := range items {
		delete(s, item)
	}
}

// Contains returns true if the item is contained in the set
func (s Set[T]) Contains(item T) bool {
	_, ok := s[item]
	return ok
}

// ContainsAll returns true if all the items are contained in the set
func (s Set[T]) ContainsAll(items Set[T]) bool {
	for item := range items {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

// Sorted returns a sorted slice based on the comparator, e.g. s.Sorted(strings.Compare)
func (s Set[T]) Sorted(comparator func(a, b T) int) []T {
	items := make([]T, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	slices.SortFunc(items, comparator)
	return items
}

// Len returns the length of the set
func (s Set[T]) Len() int {
	return len(s)
}

// Subset returns a new set with a subset of the items in this set matching the filter
func (s Set[T]) Subset(keep func(T) bool) Set[T] {
	result := NewSet[T]()
	for item := range s {
		if keep(item) {
			result[item] = struct{}{}
		}
	}
	return result
}
