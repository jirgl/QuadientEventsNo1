package core

import (
	"container/list"
)

/*
fibonacci_heap.go contains implementation of Fibonacci heap
*/

// Value interface
type Value interface {
	Value() interface{}
	Key() float64
}

//FibHeap struct
type FibHeap struct {
	roots       *list.List
	index       map[interface{}]*heapNode
	treeDegrees map[uint]*list.Element
	min         *heapNode
	num         uint
}

type heapNode struct {
	self     *list.Element
	parent   *heapNode
	children *list.List
	marked   bool
	degree   uint
	position uint
	tag      interface{}
	key      float64
	value    Value
}

//InitFibHeap creates initialized Fibonacci Heap.
func InitFibHeap() *FibHeap {
	heap := new(FibHeap)
	heap.roots = list.New()
	heap.index = make(map[interface{}]*heapNode)
	heap.treeDegrees = make(map[uint]*list.Element)
	heap.num = 0
	heap.min = nil

	return heap
}

//Len returns total number of values in heap.
func (heap *FibHeap) Len() uint {
	return heap.num
}

//Insert pushes value and key into heap.
func (heap *FibHeap) Insert(key float64, value interface{}) {
	heap.insert(value, key, nil)
}

//ExtractMin returns and extracts current minimum value and key.
func (heap *FibHeap) ExtractMin() (float64, interface{}) {
	min := heap.extractMin()

	return min.key, min.tag
}

func (heap *FibHeap) consolidate() {
	for tree := heap.roots.Front(); tree != nil; tree = tree.Next() {
		heap.treeDegrees[tree.Value.(*heapNode).position] = nil
	}

	for tree := heap.roots.Front(); tree != nil; {
		if heap.treeDegrees[tree.Value.(*heapNode).degree] == nil {
			heap.treeDegrees[tree.Value.(*heapNode).degree] = tree
			tree.Value.(*heapNode).position = tree.Value.(*heapNode).degree
			tree = tree.Next()
			continue
		}

		if heap.treeDegrees[tree.Value.(*heapNode).degree] == tree {
			tree = tree.Next()
			continue
		}

		for heap.treeDegrees[tree.Value.(*heapNode).degree] != nil {
			anotherTree := heap.treeDegrees[tree.Value.(*heapNode).degree]
			heap.treeDegrees[tree.Value.(*heapNode).degree] = nil
			if tree.Value.(*heapNode).key <= anotherTree.Value.(*heapNode).key {
				heap.roots.Remove(anotherTree)
				heap.link(tree.Value.(*heapNode), anotherTree.Value.(*heapNode))
			} else {
				heap.roots.Remove(tree)
				heap.link(anotherTree.Value.(*heapNode), tree.Value.(*heapNode))
				tree = anotherTree
			}
		}
		heap.treeDegrees[tree.Value.(*heapNode).degree] = tree
		tree.Value.(*heapNode).position = tree.Value.(*heapNode).degree
	}

	heap.resetMin()
}

func (heap *FibHeap) insert(tag interface{}, key float64, value Value) error {
	node := new(heapNode)
	node.children = list.New()
	node.tag = tag
	node.key = key
	node.value = value

	node.self = heap.roots.PushBack(node)
	heap.index[node.tag] = node
	heap.num++

	if heap.min == nil || heap.min.key > node.key {
		heap.min = node
	}

	return nil
}

func (heap *FibHeap) extractMin() *heapNode {
	min := heap.min

	children := heap.min.children
	if children != nil {
		for e := children.Front(); e != nil; e = e.Next() {
			e.Value.(*heapNode).parent = nil
			e.Value.(*heapNode).self = heap.roots.PushBack(e.Value.(*heapNode))
		}
	}

	heap.roots.Remove(heap.min.self)
	heap.treeDegrees[min.position] = nil
	delete(heap.index, heap.min.tag)
	heap.num--

	if heap.num == 0 {
		heap.min = nil
	} else {
		heap.consolidate()
	}

	return min
}

func (heap *FibHeap) link(parent, child *heapNode) {
	child.marked = false
	child.parent = parent
	child.self = parent.children.PushBack(child)
	parent.degree++
}

func (heap *FibHeap) resetMin() {
	heap.min = heap.roots.Front().Value.(*heapNode)
	for tree := heap.min.self.Next(); tree != nil; tree = tree.Next() {
		if tree.Value.(*heapNode).key < heap.min.key {
			heap.min = tree.Value.(*heapNode)
		}
	}
}
