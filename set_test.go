package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	set := New[int]()

	assert.False(t, set.Has(1))
	assert.False(t, set.Has(2))
	assert.False(t, set.Has(3))

	set.Add(1)
	set.Add(2)
	set.Add(3)

	assert.True(t, set.Has(1))
	assert.True(t, set.Has(2))
	assert.True(t, set.Has(3))
}

func TestHas(t *testing.T) {
	set := New[int]()

	assert.False(t, set.Has(1))
	assert.False(t, set.Has(2))
	assert.False(t, set.Has(3))

	set.Add(1)
	set.Add(2)
	set.Add(3)

	assert.True(t, set.Has(1))
	assert.True(t, set.Has(2))
	assert.True(t, set.Has(3))
}

func TestDelete(t *testing.T) {
	set := New[int]()

	assert.False(t, set.Has(1))
	assert.False(t, set.Has(2))
	assert.False(t, set.Has(3))

	set.Add(1)
	set.Add(2)
	set.Add(3)

	assert.True(t, set.Has(1))
	assert.True(t, set.Has(2))
	assert.True(t, set.Has(3))

	set.Delete(1)

	assert.False(t, set.Has(1))
	assert.True(t, set.Has(2))
	assert.True(t, set.Has(3))

	set.Delete(2)
	set.Delete(3)

	assert.False(t, set.Has(1))
	assert.False(t, set.Has(2))
	assert.False(t, set.Has(3))
}

func TestSize(t *testing.T) {
	set := New[int]()

	assert.Equal(t, 0, set.Size())

	set.Add(1)
	set.Add(2)
	set.Add(3)

	assert.Equal(t, 3, set.Size())

	set.Delete(1)

	assert.Equal(t, 2, set.Size())
}

func TestIsEmpty(t *testing.T) {
	set := New[int]()

	assert.True(t, set.IsEmpty())

	set.Add(1)
	set.Add(2)
	set.Add(3)

	assert.False(t, set.IsEmpty())
}

func TestString(t *testing.T) {
	set := New[int]()

	assert.Equal(t, "{ }", set.String())

	set.Add(1)
	set.Add(2)
	set.Add(3)

	assert.NotEqual(t, "", set.String())
}
