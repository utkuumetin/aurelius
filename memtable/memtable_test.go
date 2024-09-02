package memtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutShouldAddGivenKeyValuePair(t *testing.T) {
	memtable := NewMemtable()
	key := []byte("key")
	value := []byte("value")
	approximateSize := len(key) + len(value)

	memtable.Put(key, value)

	found := memtable.Find(key)
	assert.Equal(t, uint64(approximateSize), memtable.GetApproximateSize())
	assert.Equal(t, value, found)
}

func TestFindShouldReturnNilWhenKeyDoesNotExist(t *testing.T) {
	memtable := NewMemtable()
	key := []byte("key")

	value := memtable.Find(key)

	assert.Nil(t, value)
}

func TestDeleteShouldSetValueToNilWhenKeyExists(t *testing.T) {
	memtable := NewMemtable()
	key := []byte("key")
	value := []byte("value")

	memtable.Put(key, value)
	exists := memtable.Find(key)

	memtable.Delete(key)
	deleted := memtable.Find(key)

	assert.Equal(t, value, exists)
	assert.Nil(t, deleted)
}
