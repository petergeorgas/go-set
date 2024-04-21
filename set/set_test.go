package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	set := NewSet([]int{1, 2, 3, 4, 5, 6, 6, 6, 6, 7, 8}...)
	assert.Equal(t, set.Len(), 8)
	for i := 1; i <= 8; i++ {
		assert.True(t, set.Has(i))
	}
}

func TestAdd(t *testing.T) {
	set := NewSet([]int{1, 2, 3, 4, 5, 6, 6, 6, 6, 7, 8}...)
	set.Add(9)
	assert.Equal(t, set.Len(), 9)
	for i := 1; i <= 9; i++ {
		assert.True(t, set.Has(i))
	}
}

func TestRemove(t *testing.T) {
	set := NewSet([]int{1, 2, 3, 4}...)
	// Remove item that exists in set
	set.Remove(2)
	assert.Equal(t, set.Len(), 3)
	assert.False(t, set.Has(2))
	// Remove item that does not exist in set
	set.Remove(15)
	assert.Equal(t, set.Len(), 3)
	assert.False(t, set.Has(2))
}

func TestIntersection(t *testing.T) {
	set := NewSet([]int{1, 2, 3, 4}...)
	anotherSet := NewSet([]int{3, 4, 5, 6}...)

	intersect := set.Intersection(anotherSet)
	assert.Equal(t, intersect.Len(), 2)
	assert.ElementsMatch(t, intersect.Members(), []int{3, 4})
}

func TestUnion(t *testing.T) {
	set := NewSet([]int{1, 2, 3, 4}...)
	anotherSet := NewSet([]int{3, 4, 5, 6}...)

	union := set.Union(anotherSet)
	assert.Equal(t, union.Len(), 6)
	for i := 1; i <= 6; i++ {
		assert.True(t, union.Has(i))
	}
}
