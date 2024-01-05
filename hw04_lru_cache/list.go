package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Tail, Head *ListItem
	Length     int
}

func NewList() List {
	return new(list)
}

func (l *list) PushFront(v interface{}) *ListItem {
	node := &ListItem{
		Value: v,
	}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}
	l.Length++
	return node
}

func (l *list) PushBack(v interface{}) *ListItem {
	node := &ListItem{
		Value: v,
	}
	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}
	l.Length++
	return node
}

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *ListItem {
	return l.Head
}

func (l *list) Back() *ListItem {
	return l.Tail
}

func (l *list) Remove(i *ListItem) {
	if i.Prev == nil {
		l.Head = i.Next
	} else {
		i.Prev.Next = i.Next
	}

	if i.Next == nil {
		l.Tail = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}
	l.Length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	i.Prev = nil
	i.Next = l.Head
	l.Head = i
	l.Length++
}
