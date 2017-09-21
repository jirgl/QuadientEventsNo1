package core

import "testing"
import "github.com/stretchr/testify/assert"

func TestInsertIntoFibonacciHeap(t *testing.T) {
	heap := InitFibHeap()
	heap.Insert(2, "two")
	heap.Insert(1, "one")

	key, value := heap.ExtractMin()
	assert.Equal(t, float64(1), key)
	assert.Equal(t, "one", value)

	key, value = heap.ExtractMin()
	assert.Equal(t, float64(2), key)
	assert.Equal(t, "two", value)
}

func TestGetLenOfHeap(t *testing.T) {
	heap := InitFibHeap()
	heap.Insert(2, "two")
	heap.Insert(1, "one")

	assert.Equal(t, uint(2), heap.Len())
}
