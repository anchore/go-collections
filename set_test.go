package collections

import (
	"strings"
	"testing"
)

func TestNewSet(t *testing.T) {
	tests := []struct {
		name          string
		items         []int
		expectedSize  int
		shouldContain []int
	}{
		{
			name:         "empty set",
			items:        []int{},
			expectedSize: 0,
		},
		{
			name:          "with initial items",
			items:         []int{1, 2, 3},
			expectedSize:  3,
			shouldContain: []int{1, 2, 3},
		},
		{
			name:          "with duplicate items",
			items:         []int{1, 2, 2, 3, 3, 3},
			expectedSize:  3,
			shouldContain: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.items...)
			if len(s) != tt.expectedSize {
				t.Errorf("expected size %d, got %d", tt.expectedSize, len(s))
			}
			for _, item := range tt.shouldContain {
				if !s.Contains(item) {
					t.Errorf("set should contain %d", item)
				}
			}
		})
	}
}

func TestNewSet_String(t *testing.T) {
	s := NewSet("a", "b", "c")
	if len(s) != 3 {
		t.Errorf("expected size 3, got %d", len(s))
	}
	if !s.Contains("a") || !s.Contains("b") || !s.Contains("c") {
		t.Error("set should contain all initial string items")
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name          string
		initial       []int
		toAdd         []int
		expectedSize  int
		shouldContain []int
	}{
		{
			name:          "add single item",
			initial:       []int{},
			toAdd:         []int{1},
			expectedSize:  1,
			shouldContain: []int{1},
		},
		{
			name:          "add multiple items",
			initial:       []int{},
			toAdd:         []int{1, 2, 3},
			expectedSize:  3,
			shouldContain: []int{1, 2, 3},
		},
		{
			name:          "add duplicate items",
			initial:       []int{1},
			toAdd:         []int{1},
			expectedSize:  1,
			shouldContain: []int{1},
		},
		{
			name:          "add to existing set",
			initial:       []int{1, 2},
			toAdd:         []int{3, 4},
			expectedSize:  4,
			shouldContain: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.initial...)
			s.Add(tt.toAdd...)
			if len(s) != tt.expectedSize {
				t.Errorf("expected size %d, got %d", tt.expectedSize, len(s))
			}
			for _, item := range tt.shouldContain {
				if !s.Contains(item) {
					t.Errorf("set should contain %d", item)
				}
			}
		})
	}
}

func TestAddAll(t *testing.T) {
	tests := []struct {
		name          string
		initial       []int
		toAdd         []int
		expectedSize  int
		shouldContain []int
	}{
		{
			name:          "add all from empty set",
			initial:       []int{1, 2, 3},
			toAdd:         []int{},
			expectedSize:  3,
			shouldContain: []int{1, 2, 3},
		},
		{
			name:          "add all to empty set",
			initial:       []int{},
			toAdd:         []int{1, 2, 3},
			expectedSize:  3,
			shouldContain: []int{1, 2, 3},
		},
		{
			name:          "add all with no overlap",
			initial:       []int{1, 2, 3},
			toAdd:         []int{4, 5, 6},
			expectedSize:  6,
			shouldContain: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:          "add all with overlap",
			initial:       []int{1, 2, 3},
			toAdd:         []int{3, 4, 5},
			expectedSize:  5,
			shouldContain: []int{1, 2, 3, 4, 5},
		},
		{
			name:          "add all with complete overlap",
			initial:       []int{1, 2, 3},
			toAdd:         []int{1, 2, 3},
			expectedSize:  3,
			shouldContain: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.initial...)
			toAdd := NewSet(tt.toAdd...)
			s.AddAll(toAdd)
			if s.Len() != tt.expectedSize {
				t.Errorf("expected size %d, got %d", tt.expectedSize, s.Len())
			}
			for _, item := range tt.shouldContain {
				if !s.Contains(item) {
					t.Errorf("set should contain %d", item)
				}
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name             string
		initial          []int
		toRemove         []int
		expectedSize     int
		shouldContain    []int
		shouldNotContain []int
	}{
		{
			name:             "remove existing item",
			initial:          []int{1, 2, 3},
			toRemove:         []int{2},
			expectedSize:     2,
			shouldContain:    []int{1, 3},
			shouldNotContain: []int{2},
		},
		{
			name:          "remove non-existing item",
			initial:       []int{1, 2, 3},
			toRemove:      []int{4},
			expectedSize:  3,
			shouldContain: []int{1, 2, 3},
		},
		{
			name:             "remove multiple items",
			initial:          []int{1, 2, 3, 4, 5},
			toRemove:         []int{2, 4},
			expectedSize:     3,
			shouldContain:    []int{1, 3, 5},
			shouldNotContain: []int{2, 4},
		},
		{
			name:         "remove from empty set",
			initial:      []int{},
			toRemove:     []int{1},
			expectedSize: 0,
		},
		{
			name:         "remove all items",
			initial:      []int{1, 2, 3},
			toRemove:     []int{1, 2, 3},
			expectedSize: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.initial...)
			s.Remove(tt.toRemove...)
			if len(s) != tt.expectedSize {
				t.Errorf("expected size %d, got %d", tt.expectedSize, len(s))
			}
			for _, item := range tt.shouldContain {
				if !s.Contains(item) {
					t.Errorf("set should contain %d", item)
				}
			}
			for _, item := range tt.shouldNotContain {
				if s.Contains(item) {
					t.Errorf("set should not contain %d", item)
				}
			}
		})
	}
}

func TestRemoveAll(t *testing.T) {
	tests := []struct {
		name             string
		initial          []int
		toRemove         []int
		expectedSize     int
		shouldContain    []int
		shouldNotContain []int
	}{
		{
			name:          "remove from empty set",
			initial:       []int{},
			toRemove:      []int{1, 2, 3},
			expectedSize:  0,
			shouldContain: []int{},
		},
		{
			name:          "remove empty set",
			initial:       []int{1, 2, 3},
			toRemove:      []int{},
			expectedSize:  3,
			shouldContain: []int{1, 2, 3},
		},
		{
			name:             "remove some items",
			initial:          []int{1, 2, 3, 4, 5},
			toRemove:         []int{2, 4},
			expectedSize:     3,
			shouldContain:    []int{1, 3, 5},
			shouldNotContain: []int{2, 4},
		},
		{
			name:             "remove all items",
			initial:          []int{1, 2, 3},
			toRemove:         []int{1, 2, 3, 4, 5},
			expectedSize:     0,
			shouldNotContain: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.initial...)
			toRemove := NewSet(tt.toRemove...)
			s.RemoveAll(toRemove)
			if s.Len() != tt.expectedSize {
				t.Errorf("expected size %d, got %d", tt.expectedSize, s.Len())
			}
			for _, item := range tt.shouldContain {
				if !s.Contains(item) {
					t.Errorf("set should contain %d", item)
				}
			}
			for _, item := range tt.shouldNotContain {
				if s.Contains(item) {
					t.Errorf("set should not contain %d", item)
				}
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name          string
		initial       []int
		checkItem     int
		shouldContain bool
	}{
		{
			name:          "contains existing item",
			initial:       []int{1, 2, 3},
			checkItem:     2,
			shouldContain: true,
		},
		{
			name:          "does not contain non-existing item",
			initial:       []int{1, 2, 3},
			checkItem:     4,
			shouldContain: false,
		},
		{
			name:          "empty set contains nothing",
			initial:       []int{},
			checkItem:     1,
			shouldContain: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.initial...)
			result := s.Contains(tt.checkItem)
			if result != tt.shouldContain {
				t.Errorf("expected Contains(%d) to be %v, got %v", tt.checkItem, tt.shouldContain, result)
			}
		})
	}
}

func TestContains_AfterAdd(t *testing.T) {
	s := NewSet[int]()
	s.Add(5)
	if !s.Contains(5) {
		t.Error("set should contain item after adding")
	}
}

func TestContains_AfterRemove(t *testing.T) {
	s := NewSet(1, 2, 3)
	s.Remove(2)
	if s.Contains(2) {
		t.Error("set should not contain item after removing")
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		name     string
		set      []int
		check    []int
		expected bool
	}{
		{
			name:     "both empty",
			set:      []int{},
			check:    []int{},
			expected: true,
		},
		{
			name:     "check empty set",
			set:      []int{1, 2, 3},
			check:    []int{},
			expected: true,
		},
		{
			name:     "contains all items",
			set:      []int{1, 2, 3, 4, 5},
			check:    []int{2, 3, 4},
			expected: true,
		},
		{
			name:     "contains exact items",
			set:      []int{1, 2, 3},
			check:    []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "does not contain all items",
			set:      []int{1, 2, 3},
			check:    []int{2, 3, 4},
			expected: false,
		},
		{
			name:     "empty set does not contain items",
			set:      []int{},
			check:    []int{1, 2, 3},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.set...)
			check := NewSet(tt.check...)
			result := s.ContainsAll(check)
			if result != tt.expected {
				t.Errorf("expected ContainsAll() = %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSet_Bool(t *testing.T) {
	s := NewSet(true, false)
	if len(s) != 2 {
		t.Errorf("expected size 2, got %d", len(s))
	}
	if !s.Contains(true) || !s.Contains(false) {
		t.Error("set should contain both boolean values")
	}
}

func TestSet_Struct(t *testing.T) {
	type Point struct {
		X, Y int
	}
	s := NewSet(Point{1, 2}, Point{3, 4})
	if len(s) != 2 {
		t.Errorf("expected size 2, got %d", len(s))
	}
	if !s.Contains(Point{1, 2}) {
		t.Error("set should contain Point{1, 2}")
	}
}

func TestSetOperations(t *testing.T) {
	tests := []struct {
		name             string
		operations       func() Set[int]
		expectedSize     int
		shouldContain    []int
		shouldNotContain []int
	}{
		{
			name: "add and remove sequence",
			operations: func() Set[int] {
				s := NewSet[int]()
				s.Add(1)
				s.Add(2)
				s.Remove(1)
				s.Add(3)
				return s
			},
			expectedSize:     2,
			shouldContain:    []int{2, 3},
			shouldNotContain: []int{1},
		},
		{
			name: "concurrent modifications",
			operations: func() Set[int] {
				s := NewSet(1, 2, 3)
				s.Add(4, 5)
				s.Remove(1, 2)
				s.Add(6)
				return s
			},
			expectedSize:     4,
			shouldContain:    []int{3, 4, 5, 6},
			shouldNotContain: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.operations()
			if len(s) != tt.expectedSize {
				t.Errorf("expected size %d, got %d", tt.expectedSize, len(s))
			}
			for _, item := range tt.shouldContain {
				if !s.Contains(item) {
					t.Errorf("set should contain %d", item)
				}
			}
			for _, item := range tt.shouldNotContain {
				if s.Contains(item) {
					t.Errorf("set should not contain %d", item)
				}
			}
		})
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name     string
		items    []int
		expected int
	}{
		{
			name:     "empty set",
			items:    []int{},
			expected: 0,
		},
		{
			name:     "single item",
			items:    []int{1},
			expected: 1,
		},
		{
			name:     "multiple items",
			items:    []int{1, 2, 3, 4, 5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.items...)
			if s.Len() != tt.expected {
				t.Errorf("expected Len() = %d, got %d", tt.expected, s.Len())
			}
		})
	}
}

func TestSubset(t *testing.T) {
	tests := []struct {
		name          string
		items         []int
		predicate     func(int) bool
		expectedSize  int
		shouldContain []int
	}{
		{
			name:          "empty set",
			items:         []int{},
			predicate:     func(n int) bool { return n > 0 },
			expectedSize:  0,
			shouldContain: []int{},
		},
		{
			name:          "filter evens",
			items:         []int{1, 2, 3, 4, 5, 6},
			predicate:     func(n int) bool { return n%2 == 0 },
			expectedSize:  3,
			shouldContain: []int{2, 4, 6},
		},
		{
			name:          "filter odds",
			items:         []int{1, 2, 3, 4, 5, 6},
			predicate:     func(n int) bool { return n%2 != 0 },
			expectedSize:  3,
			shouldContain: []int{1, 3, 5},
		},
		{
			name:          "filter greater than",
			items:         []int{1, 2, 3, 4, 5},
			predicate:     func(n int) bool { return n > 3 },
			expectedSize:  2,
			shouldContain: []int{4, 5},
		},
		{
			name:          "no matches",
			items:         []int{1, 2, 3},
			predicate:     func(n int) bool { return n > 10 },
			expectedSize:  0,
			shouldContain: []int{},
		},
		{
			name:          "all match",
			items:         []int{1, 2, 3},
			predicate:     func(n int) bool { return n > 0 },
			expectedSize:  3,
			shouldContain: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.items...)
			result := s.Subset(tt.predicate)
			if result.Len() != tt.expectedSize {
				t.Errorf("expected size %d, got %d", tt.expectedSize, result.Len())
			}
			for _, item := range tt.shouldContain {
				if !result.Contains(item) {
					t.Errorf("subset should contain %d", item)
				}
			}
		})
	}
}

func TestSorted(t *testing.T) {
	tests := []struct {
		name     string
		items    []string
		expected []string
	}{
		{
			name:     "empty set",
			items:    []string{},
			expected: []string{},
		},
		{
			name:     "single item",
			items:    []string{"apple"},
			expected: []string{"apple"},
		},
		{
			name:     "already sorted",
			items:    []string{"apple", "banana", "cherry", "date", "elderberry"},
			expected: []string{"apple", "banana", "cherry", "date", "elderberry"},
		},
		{
			name:     "reverse sorted",
			items:    []string{"elderberry", "date", "cherry", "banana", "apple"},
			expected: []string{"apple", "banana", "cherry", "date", "elderberry"},
		},
		{
			name:     "random order",
			items:    []string{"grape", "apple", "fig", "apple", "kiwi", "banana", "cherry"},
			expected: []string{"apple", "banana", "cherry", "fig", "grape", "kiwi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.items...)
			result := s.Sorted(strings.Compare)
			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
			}
			for i, v := range tt.expected {
				if result[i] != v {
					t.Errorf("at index %d: expected %q, got %q", i, v, result[i])
				}
			}
		})
	}
}
