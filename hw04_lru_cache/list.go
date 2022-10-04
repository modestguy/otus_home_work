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
	size int
	head *ListItem
	tail *ListItem
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	l.Remove(i)
	l.PushFront(i.Value)
}

func (l *list) Remove(i *ListItem) {
	item := l.head
	for {
		if item == i && item.Value == i.Value || item == nil {
			break
		}
		item = item.Next
	}
	if item == nil {
		return
	}

	prev := item.Prev
	next := item.Next
	if prev != nil {
		prev.Next = next
	}

	if next != nil {
		next.Prev = prev
	}

	l.size--
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	item.Next = l.head
	item.Prev = nil
	if l.head != nil {
		l.head.Prev = item
	}
	l.head = item
	if l.tail == nil {
		l.tail = item
	}
	l.size++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	item.Next = nil
	item.Prev = l.tail
	if l.tail != nil {
		l.tail.Next = item
	}
	l.tail = item
	if l.head == nil {
		l.head = item
	}
	l.size++
	return item
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func NewList() List {
	return new(list)
}
