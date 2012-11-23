// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//Modifications; Mortdeus(mortdeus.mit-license.org)

// Package list implements a doubly linked list.
//
// To iterate over a list (where l is a *List):
//	for e := l.Front(); e != nil; e = e.Next() {
//		// do something with e.Value
//	}
//
package gocos2d

// Child is an Child in the linked list.
type Child struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// The front of the list has prev = nil, and the back has next = nil.
	next, prev *Child

	// The list to which this Child belongs.
	list *ChildList

	Value Node_
}

// Next returns the next list Child or nil.
func (e *Child) Next() *Child { return e.next }

// Prev returns the previous list Child or nil.
func (e *Child) Prev() *Child { return e.prev }

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type ChildList struct {
	dict        map[string]Node_
	front, back *Child
	len         int
}

func (l *ChildList) Lookup(tag string) Node_ {
	return l.dict[tag]
}

// Init initializes or clears a List.
func (l *ChildList) Init() *ChildList {
	l.front = nil
	l.back = nil
	l.len = 0
	l.dict = make(map[string]Node_)
	return l
}

// Front returns the first Child in the list.
func (l *ChildList) Front() *Child { return l.front }

// Back returns the last Child in the list.
func (l *ChildList) Back() *Child { return l.back }

// Remove removes the Child from the list
// and returns its Value.
func (l *ChildList) Remove(e *Child) Node_ {
	l.remove(e)
	e.list = nil // do what remove does not
	delete(l.dict, e.Value.Tag())
	return e.Value
}

// remove the Child from the list, but do not clear the Child's list field.
// This is so that other List methods may use remove when relocating Elements
// without needing to restore the list field.
func (l *ChildList) remove(e *Child) {
	if e.list != l {
		return
	}
	if e.prev == nil {
		l.front = e.next
	} else {
		e.prev.next = e.next
	}
	if e.next == nil {
		l.back = e.prev
	} else {
		e.next.prev = e.prev
	}

	e.prev = nil
	e.next = nil
	l.len--
}

func (l *ChildList) insertBefore(e *Child, mark *Child) {
	if mark.prev == nil {
		// new front of the list
		l.front = e
	} else {
		mark.prev.next = e
	}
	e.prev = mark.prev
	mark.prev = e
	e.next = mark
	l.len++
}

func (l *ChildList) insertAfter(e *Child, mark *Child) {
	if mark.next == nil {
		// new back of the list
		l.back = e
	} else {
		mark.next.prev = e
	}
	e.next = mark.next
	mark.next = e
	e.prev = mark
	l.len++
}

func (l *ChildList) insertFront(e *Child) {
	if l.front == nil {
		// empty list
		l.front, l.back = e, e
		e.prev, e.next = nil, nil
		l.len = 1
		return
	}
	l.insertBefore(e, l.front)
}

func (l *ChildList) insertBack(e *Child) {
	if l.back == nil {
		// empty list
		l.front, l.back = e, e
		e.prev, e.next = nil, nil
		l.len = 1
		return
	}
	l.insertAfter(e, l.back)
}

// PushFront inserts the value at the front of the list and returns a new Child containing the value.
func (l *ChildList) PushFront(value Node_) *Child {
	e := &Child{nil, nil, l, value}
	l.dict[value.Tag()] = value
	l.insertFront(e)
	return e
}

// PushBack inserts the value at the back of the list and returns a new Child containing the value.
func (l *ChildList) PushBack(value Node_) *Child {
	e := &Child{nil, nil, l, value}
	l.dict[value.Tag()] = value
	l.insertBack(e)
	return e
}

// InsertBefore inserts the value immediately before mark and returns a new Child containing the value.
func (l *ChildList) InsertBefore(value Node_, mark *Child) *Child {
	if mark.list != l {
		return nil
	}
	e := &Child{nil, nil, l, value}
	l.dict[value.Tag()] = value
	l.insertBefore(e, mark)
	return e
}

// InsertAfter inserts the value immediately after mark and returns a new Child containing the value.
func (l *ChildList) InsertAfter(value Node_, mark *Child) *Child {
	if mark.list != l {
		return nil
	}
	e := &Child{nil, nil, l, value}
	l.dict[value.Tag()] = value
	l.insertAfter(e, mark)
	return e
}

// MoveToFront moves the Child to the front of the list.
func (l *ChildList) MoveToFront(e *Child) {
	if e.list != l || l.front == e {
		return
	}
	l.remove(e)
	l.insertFront(e)
}

// MoveToBack moves the Child to the back of the list.
func (l *ChildList) MoveToBack(e *Child) {
	if e.list != l || l.back == e {
		return
	}
	l.remove(e)
	l.insertBack(e)
}

// Len returns the number of elements in the list.
func (l *ChildList) Len() int { return l.len }

// PushBackList inserts each Child of ol at the back of the list.
func (l *ChildList) PushBackList(ol *ChildList) {
	last := ol.Back()
	for e := ol.Front(); e != nil; e = e.Next() {
		l.PushBack(e.Value)
		if e == last {
			break
		}
	}
}

// PushFrontList inserts each Child of ol at the front of the list. The ordering of the passed list is preserved.
func (l *ChildList) PushFrontList(ol *ChildList) {
	first := ol.Front()
	for e := ol.Back(); e != nil; e = e.Prev() {
		l.PushFront(e.Value)
		if e == first {
			break
		}
	}
}
