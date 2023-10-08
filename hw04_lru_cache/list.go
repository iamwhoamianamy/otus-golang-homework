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
	head     *ListItem
	tail     *ListItem
	length   int
	allItems map[*ListItem]bool
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newHead := ListItem{
		Value: v,
		Prev:  nil,
		Next:  l.head,
	}

	if l.head != nil {
		l.head.Prev = &newHead
	}

	l.head = &newHead
	l.allItems[&newHead] = true

	if l.tail == nil {
		l.tail = &newHead
	}

	l.length++
	return &newHead
}

func (l *list) PushBack(v interface{}) *ListItem {
	newTail := ListItem{
		Value: v,
		Prev:  l.tail,
		Next:  nil,
	}

	if l.tail != nil {
		l.tail.Next = &newTail
	}

	l.tail = &newTail
	l.allItems[&newTail] = true

	if l.head == nil {
		l.head = &newTail
	}

	l.length++
	return &newTail
}

func (l *list) detach(i *ListItem) {
	if i == l.head {
		l.head = l.head.Next
	}

	if i == l.tail {
		l.tail = l.tail.Prev
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	i.Next = nil
	i.Prev = nil
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	_, ok := l.allItems[i]

	if !ok {
		return
	}

	l.detach(i)

	delete(l.allItems, i)
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.detach(i)

	if l.head != nil {
		l.head.Prev = i
	}

	i.Next = l.head
	l.head = i
}

func NewList() List {
	return &list{
		head:     nil,
		tail:     nil,
		length:   0,
		allItems: make(map[*ListItem]bool),
	}
}
