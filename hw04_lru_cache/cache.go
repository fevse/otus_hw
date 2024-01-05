package hw04lrucache

type Key string

type Item struct {
	Key   Key
	Value interface{}
}

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

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if i, ok := l.items[key]; ok {
		l.queue.MoveToFront(i)
		i.Value.(*Item).Value = value
		return true
	}

	if l.queue.Len() == l.capacity {
		if i := l.queue.Back(); i != nil {
			item := i.Value.(*Item)
			l.queue.Remove(i)
			delete(l.items, item.Key)
		}
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	i := l.queue.PushFront(item)
	l.items[item.Key] = i

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if i, ok := l.items[key]; ok {
		l.queue.MoveToFront(i)
		return i.Value.(*Item).Value, ok
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l = &lruCache{
		capacity: l.capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, l.capacity),
	}
}
