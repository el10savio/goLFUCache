package lfu

import (
	"errors"

	linkedList "github.com/el10savio/goLFUCache/linkedList"
)

// LFU ...
type LFU struct {
	Capacity      uint
	hashTable     map[int]int
	accessList    map[*linkedList.Node][]CacheEntry
	frequencyList linkedList.DoublyLinkedList
}

// CacheEntry ...
type CacheEntry struct {
	Key   int
	Value int
}

// InitLFUCache ...
func InitLFUCache(capacity uint) (*LFU, error) {
	if capacity == 0 {
		return nil, errors.New("cannot create lfu cache with 0 capacity")
	}

	lfu := &LFU{
		Capacity:      capacity,
		hashTable:     make(map[int]int),
		accessList:    make(map[*linkedList.Node][]CacheEntry),
		frequencyList: *linkedList.InitDoublyLinkedList(),
	}

	return lfu, nil
}

// ClearLFUCache ...
func (lfu *LFU) ClearLFUCache() {
	lfu, _ = InitLFUCache(lfu.Capacity)
}

// NewEntry ...
func (lfu *LFU) NewEntry(key, value int) error {
	if _, ok := lfu.hashTable[key]; ok {
		return errors.New("key already present in cache")
	}

	// Add To HashTable
	lfu.hashTable[key] = value

	// Add "1" To frequencyList
	lfu.frequencyList.Append(1)

	// Add Entry To 1's AccessList
	err := lfu.AddEntryToAccessList(1, CacheEntry{key, value})
	if err != nil {
		return err
	}

	return nil
}

// AddEntryToAccessList ...
func (lfu *LFU) AddEntryToAccessList(nodeValue int, cacheEntry CacheEntry) error {
	frequencyNode, err := lfu.frequencyList.FindElement(nodeValue)
	if err != nil {
		return err
	}

	lfu.accessList[frequencyNode] = append(lfu.accessList[frequencyNode], cacheEntry)

	return nil
}

// GetEntry ...
func (lfu *LFU) GetEntry(key int) (int, error) {
	if _, ok := lfu.hashTable[key]; !ok {
		return 0, errors.New("key is not present in cache")
	}

	// Update frequencyList
	node, listIndex := lfu.FindNode(key)
	next := node.Next

	if next == nil {
		lfu.frequencyList.Append(node.Value + 1)
	} else if next.Value != node.Value+1 {
		lfu.frequencyList.InsertBetween(node.Value+1, node.Value, next.Value)
	}

	err := lfu.AddEntryToAccessList(node.Value+1, CacheEntry{key, lfu.hashTable[key]})
	if err != nil {
		return 0, err
	}

	// Delete Old frequencyNode's Entry
	lfu.accessList[node] = append(lfu.accessList[node][:listIndex], lfu.accessList[node][listIndex+1:]...)

	return lfu.hashTable[key], nil
}

// FindNode ...
func (lfu *LFU) FindNode(key int) (*linkedList.Node, int) {
	head := lfu.frequencyList.Head

	for head != nil {
		for index, list := range lfu.accessList[head] {
			if key == list.Key {
				return head, index
			}
			head = head.Next
		}
	}

	return nil, 0
}
