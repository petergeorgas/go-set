package set

type Set[T comparable] interface {
	Add(item T)                     // Insert an item into the set
	Has(item T) bool                // Whether or not an item exists in the set
	Remove(item T)                  // Removes an item from the set
	Union(set Set[T]) Set[T]        // Union of two sets (all elements of both sets combined)
	Intersection(set Set[T]) Set[T] // Intersection of two sets (only common elements of both sets)
	Len() int                       // Number of members in the set
	Members() []T                   // Utility for converting to slice
}

type set[T comparable] struct {
	vals map[T]struct{}
}

func (s *set[T]) Add(item T) {
	s.vals[item] = struct{}{}
}

func (s *set[T]) Has(item T) bool {
	_, ok := s.vals[item]
	return ok
}

func (s *set[T]) Remove(item T) {
	_, ok := s.vals[item]
	if !ok {
		return
	}

	delete(s.vals, item)
}

func (s *set[T]) Union(set Set[T]) Set[T] {
	vals := append(s.Members(), set.Members()...)
	return NewSet(vals...)
}

func (s *set[T]) Intersection(set Set[T]) Set[T] {
	vals := []T{}
	for _, item := range set.Members() {
		if !s.Has(item) {
			continue
		}
		vals = append(vals, item)
	}

	return NewSet(vals...)
}

func (s *set[T]) Len() int { return len(s.vals) }
func (s *set[T]) Members() []T {
	vals := make([]T, s.Len())
	idx := 0
	for item := range s.vals {
		vals[idx] = item
		idx++
	}

	return vals
}

func NewSet[T comparable](vals ...T) Set[T] {
	values := make(map[T]struct{}, len(vals))
	for _, val := range vals {
		values[val] = struct{}{}
	}

	return &set[T]{
		vals: values,
	}
}
