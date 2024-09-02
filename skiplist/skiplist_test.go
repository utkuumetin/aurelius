package skiplist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutShouldCreateNewNode(t *testing.T) {
	key := []byte("key")
	value := []byte("value")
	skiplist := NewSkipList()

	skiplist.Put(key, value)

	node := skiplist.Find(key)
	assert.Equal(t, node.Key, key)
	assert.Equal(t, node.Value, value)
}

func TestPutShouldChangeExistingNodeValueWhenGivenKeyExists(t *testing.T) {
	key := []byte("key")
	value := []byte("value")
	skiplist := NewSkipList()

	skiplist.Put(key, value)

	appendix := []byte(" is a new value")
	value = append(value, appendix...)

	skiplist.Put(key, value)

	node := skiplist.Find(key)
	assert.Equal(t, node.Key, key)
	assert.Equal(t, node.Value, value)
}

func TestFindGreaterOrEqualShouldReturnNilValueWhenGivenKeyOrGreaterKeyDoNotExist(t *testing.T) {
	skipList := NewSkipList()
	key := []byte("key")
	value := []byte("value")

	skipList.Put(key, value)

	appendix := []byte(" is greater")
	key = append(key, appendix...)

	node, _ := skipList.FindGreaterOrEqual(key)
	assert.Nil(t, node)
}
