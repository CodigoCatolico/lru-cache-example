package main

import (
	"sync"
)

// lruCache is a least recently used cache implemented with a
// double linked list.
type lruCache struct {
	cacheMap map[string]*lruItem
	maxSize  int
	size     int
	head     *lruItem
	tail     *lruItem
	mx       sync.Mutex
}

// lruItem is an item of the cache linked list.
type lruItem struct {
	next *lruItem
	prev *lruItem
	key  string
	data interface{}
}

// newLRUCache creates a new LRU Cache and returns its address.
func newLRUCache(maxSize int) *lruCache {
	return &lruCache{
		cacheMap: make(map[string]*lruItem),
		maxSize:  maxSize,
	}
}

// put puts a new value in the cache with the given key.
// If the entry already exists it overwrites its value,
// if not it is created. When put, the item is always moved
// to the head.
func (l *lruCache) put(key string, data interface{}) {
	l.mx.Lock()
	defer l.mx.Unlock()

	// Tries to get an item from the cache.
	if existingItem, hit := l.cacheMap[key]; hit {
		// If it was a cache hit, updates the cache value
		// and returns.
		existingItem.data = data
		// When an item's value is set it is moved to the head.
		l.setHead(existingItem)
		return
	}

	// If it wasn't a hit, creates a new item to be added.
	l.size++
	newItem := &lruItem{key: key, data: data}
	l.cacheMap[newItem.key] = newItem

	// When an item is created it is the new head.
	l.setHead(newItem)

	// If the current size of the cache is bigger than maxSize
	// remove the last item from the cache (which is the least
	// recently used).
	if l.size > l.maxSize {
		l.removeItem(l.tail)
	}
}

// remove removes an item from the cache based on the key
// informed.
func (l *lruCache) remove(key string) {
	l.mx.Lock()
	defer l.mx.Unlock()

	item, hit := l.cacheMap[key]
	if !hit {
		return
	}

	l.removeItem(item)
}

// removeItem removes an item from the cache.
func (l *lruCache) removeItem(item *lruItem) {
	delete(l.cacheMap, (*item).key)
	l.unlink(item)
	l.size--
}

// retrieve retrieves the value related to the key informed.
// The second argument returned is a boolean to inform if the
// key hit the cache or if it missed.
func (l *lruCache) retrieve(key string) (interface{}, bool) {
	l.mx.Lock()
	defer l.mx.Unlock()

	item, hit := l.cacheMap[key]
	if !hit {
		return nil, false
	}

	l.setHead(item)

	return item.data, true
}

// setHead sets the linked list head to requested item.
func (l *lruCache) setHead(item *lruItem) {
	if l.head == item {
		return
	}

	l.unlink(item)
	item.next = l.head

	if l.head != nil {
		l.head.prev = item
	}

	l.head = item

	if l.tail == nil {
		l.tail = item
	}
}

// unlink removes a link from the linked list.
func (l *lruCache) unlink(item *lruItem) {
	if l.head == item {
		l.head = item.next
	}

	if l.tail == item {
		l.tail = item.prev
	}

	if item.next != nil {
		item.next.prev = item.prev
	}

	if item.prev != nil {
		item.prev.next = item.next
	}
}
