package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mutex    sync.RWMutex
	capacity int
	queue    List
	items    map[Key]*ListItem
	keys     map[*ListItem]Key
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.RLock()
	i, ok := c.items[key]
	c.mutex.RUnlock()

	if ok {
		c.mutex.Lock()

		i.Value = value
		c.queue.MoveToFront(i)

		c.mutex.Unlock()
	} else {
		c.mutex.RLock()

		if c.queue.Len() == c.capacity {
			itemToRemove := c.queue.Back()
			keyToRemove := c.keys[itemToRemove]
			c.mutex.RUnlock()

			c.mutex.Lock()

			c.queue.Remove(itemToRemove)
			delete(c.items, keyToRemove)
			delete(c.keys, itemToRemove)

			c.mutex.Unlock()
		} else {
			c.mutex.RUnlock()
		}

		c.mutex.Lock()

		i = c.queue.PushFront(value)
		c.items[key] = i
		c.keys[i] = key

		c.mutex.Unlock()
	}

	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.RLock()
	i, ok := c.items[key]
	c.mutex.RUnlock()

	if ok {
		c.mutex.Lock()

		c.queue.MoveToFront(i)

		c.mutex.Unlock()

		return i.Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.mutex.RLock()
	length := c.queue.Len()
	c.mutex.RUnlock()

	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Чтобы было за O(1) надо чтобы у списка был Clear() за O(1)
	// (который будет просто отцеплять head и tail и выставять len = 0)
	for i := 0; i < length; i++ {
		c.queue.Remove(c.queue.Front())
	}

	c.items = make(map[Key]*ListItem)
	c.keys = make(map[*ListItem]Key)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		keys:     make(map[*ListItem]Key, capacity),
	}
}
