package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	data map[T]struct{}
}

// New constructs and returns an empty set.
// time-complexity: O(1)
func New[T comparable]() Set[T] {
	return Set[T]{make(map[T]struct{})}
}

// Add adds a new elements to the set.
// time-complexity: O(1)
func (s *Set[T]) Add(val T) {
	s.data[val] = struct{}{}
}

// Has returns true if the set has the given val.
// time-complexity: O(1)
func (s *Set[T]) Has(val T) bool {
	_, ok := s.data[val]
	return ok
}

// Delete removes the val from the set.
// time-complexity: O(1)
func (s *Set[T]) Delete(val T) {
	delete(s.data, val)
}

// Size returns the number of the elements in the set.
// time-complexity: O(1)
func (s *Set[T]) Size() int {
	return len(s.data)
}

// IsEmpty returns true if the set doesn't contain any elements.
// time-complexity: O(1)
func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

// String returns the string representation of the set.
// time-complexity: O(n)
func (s *Set[T]) String() string {
	var b strings.Builder

	b.WriteString("{ ")

	for k := range s.data {
		b.WriteString(fmt.Sprint(k))
		b.WriteString(" ")
	}

	b.WriteString("}")

	return b.String()
}
