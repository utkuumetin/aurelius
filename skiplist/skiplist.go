package skiplist

import (
	"bytes"
	"math/rand/v2"
)

const (
	DefaultMaxLevel    int     = 16
	DefaultProbability float64 = 0.5
)

type SkipList struct {
	Head   *SkipListNode
	Height int
}

type SkipListNode struct {
	Key   []byte
	Value []byte
	Next  [DefaultMaxLevel]*SkipListNode
}

func NewSkipList() *SkipList {
	return &SkipList{
		Head:   &SkipListNode{},
		Height: 1,
	}
}

func (sl *SkipList) Put(key, value []byte) {
	node, previous := sl.FindGreaterOrEqual(key)
	if node != nil && bytes.Equal(key, node.Key) {
		node.Value = value
		return
	}

	node = &SkipListNode{Key: key, Value: value}
	height := generateRandomHeight()
	if height > sl.Height {
		for level := sl.Height; level < height; level++ {
			previous[level] = sl.Head
		}

		sl.Height = height
	}

	for level := height - 1; level >= 0; level-- {
		node.Next[level] = previous[level].Next[level]
		previous[level].Next[level] = node
	}
}

func (sl *SkipList) Find(key []byte) *SkipListNode {
	node, _ := sl.FindGreaterOrEqual(key)
	if node != nil && bytes.Equal(key, node.Key) {
		return node
	}

	return nil
}

func (sl *SkipList) FindGreaterOrEqual(key []byte) (*SkipListNode, [DefaultMaxLevel]*SkipListNode) {
	var current *SkipListNode
	var previous [DefaultMaxLevel]*SkipListNode

	current = sl.Head
	for level := sl.Height - 1; level >= 0; level-- {
		for current.Next[level] != nil && bytes.Compare(key, current.Next[level].Key) > 0 {
			current = current.Next[level]
		}

		previous[level] = current
	}

	return current.Next[0], previous
}

func generateRandomHeight() int {
	level := 1
	for level < DefaultMaxLevel && rand.Float64() < DefaultProbability {
		level++
	}

	return level
}
