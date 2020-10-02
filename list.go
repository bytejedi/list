// Package list provides a simple wrapper to go's container/list with the addition af a mutex
package list

import (
	"container/list"
	"sync"
)

// List represents a thread safe doubly linked list
type List struct {
	list *list.List
	mu   sync.RWMutex
}

func New() *List {
	return &List{list: list.New()}
}

func (l *List) Back() *list.Element {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.Back()
}

func (l *List) Front() *list.Element {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.Front()
}

func (l *List) Init() *List {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.Init()
	return l
}

func (l *List) InsertAfter(v interface{}, mark *list.Element) *list.Element {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.InsertAfter(v, mark)
}

func (l *List) InsertBefore(v interface{}, mark *list.Element) *list.Element {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.InsertBefore(v, mark)
}

func (l *List) Len() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list.Len()
}

func (l *List) MoveAfter(e, mark *list.Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.MoveAfter(e, mark)
}

func (l *List) MoveBefore(e, mark *list.Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.MoveBefore(e, mark)
}

func (l *List) MoveToBack(e *list.Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.MoveToBack(e)
}

func (l *List) MoveToFront(e *list.Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.MoveToFront(e)
}

func (l *List) PushBack(v interface{}) *list.Element {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.PushBack(v)
}

func (l *List) PushBackList(other *List) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.PushBackList(other.list)
}

func (l *List) PushFront(v interface{}) *list.Element {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.PushFront(v)
}

func (l *List) PushFrontList(other *List) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list.PushFrontList(other.list)
}

func (l *List) Remove(e *list.Element) interface{} {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list.Remove(e)
}
