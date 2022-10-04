package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Clear() {
	for _, elem := range c.items {
		c.queue.Remove(elem)
	}
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	element, exists := c.items[key]
	if !exists {
		return nil, false
	}
	c.queue.MoveToFront(element)
	return element.Value.(*cacheItem).value, true
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if element, exists := c.items[key]; exists {
		c.queue.MoveToFront(element)
		element.Value.(*cacheItem).value = value
		return true
	}

	if c.queue.Len() == c.capacity {
		c.purge()
	}

	item := &cacheItem{
		key:   key,
		value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.key] = element

	return false
}

func mapKey(m map[Key]*ListItem, value interface{}) (key Key, ok bool) {
	for k, v := range m {
		if v.Value == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func (c *lruCache) purge() {
	if element := c.queue.Back(); element != nil {
		key, ok := mapKey(c.items, element.Value)
		if ok {
			c.queue.Remove(element)
			delete(c.items, key)
		}
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
