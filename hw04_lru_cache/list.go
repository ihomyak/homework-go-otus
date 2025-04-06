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
	len   int
	start *ListItem
	end   *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.start
}

func (l *list) Back() *ListItem {
	return l.end
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := ListItem{Value: v, Next: l.start, Prev: nil}
	if l.start != nil {
		l.start.Prev = &item
	}
	if l.end == nil {
		l.end = &item
	}
	l.start = &item
	l.len++
	return l.start
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := ListItem{Value: v, Next: nil, Prev: l.end}
	if l.end != nil {
		l.end.Next = &item
	}
	if l.start == nil {
		l.start = &item
	}
	l.end = &item
	l.len++
	return l.end
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || l.start == i {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.end = i.Prev
	}

	i.Next = l.start
	i.Prev = nil
	l.start.Prev = i
	l.start = i
}
