package memtable

import (
	"sync"

	"github.com/utkuumetin/aurelius/skiplist"
)

type Memtable struct {
	mutex           sync.RWMutex
	list            *skiplist.SkipList
	approximateSize uint64
}

func NewMemtable() *Memtable {
	return &Memtable{
		list: skiplist.NewSkipList(),
	}
}

func (memtable *Memtable) Put(key, value []byte) {
	memtable.mutex.Lock()
	defer memtable.mutex.Unlock()

	estimatedSize := len(key) + len(value)

	memtable.list.Put(key, value)
	memtable.approximateSize = memtable.approximateSize + uint64(estimatedSize)
}

func (memtable *Memtable) Find(key []byte) []byte {
	memtable.mutex.RLock()
	defer memtable.mutex.RUnlock()

	node := memtable.list.Find(key)
	if node != nil && node.Value != nil {
		return node.Value
	}

	return nil
}

func (memtable *Memtable) Delete(key []byte) {
	memtable.mutex.Lock()
	defer memtable.mutex.Unlock()

	memtable.list.Put(key, nil)
}

func (memtable *Memtable) GetApproximateSize() uint64 {
	return memtable.approximateSize
}
