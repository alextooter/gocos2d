package gocos2d

type Child struct {
	next, prev *Child
	list       *ChildList
	Node       Node_
}

func (e *Child) Next() *Child { return e.next }
func (e *Child) Prev() *Child { return e.prev }

type ChildList struct {
	dict        map[string]*Child
	front, back *Child
	len         int
}

func (l *ChildList) Lookup(tag string) *Child {
	return l.dict[tag]
}

func (l *ChildList) Init() *ChildList {
	l.front = nil
	l.back = nil
	l.len = 0
	l.dict = make(map[string]*Child)
	return l
}

func (l *ChildList) Front() *Child { return l.front }

func (l *ChildList) Back() *Child { return l.back }

func (l *ChildList) Remove(e *Child) Node_ {
	l.remove(e)
	e.list = nil // do what remove does not
	delete(l.dict, e.Node.Tag())
	return e.Node
}

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

func (l *ChildList) PushFront(node Node_) *Child {
	e := &Child{nil, nil, l, node}
	l.dict[node.Tag()] = e
	l.insertFront(e)
	return e
}

func (l *ChildList) PushBack(node Node_) *Child {
	e := &Child{nil, nil, l, node}
	l.dict[node.Tag()] = e
	l.insertBack(e)
	return e
}

func (l *ChildList) InsertBefore(node Node_, mark *Child) *Child {
	if mark.list != l {
		return nil
	}
	e := &Child{nil, nil, l, node}
	l.dict[node.Tag()] = e
	l.insertBefore(e, mark)
	return e
}

func (l *ChildList) InsertAfter(node Node_, mark *Child) *Child {
	if mark.list != l {
		return nil
	}
	e := &Child{nil, nil, l, node}
	l.dict[node.Tag()] = e
	l.insertAfter(e, mark)
	return e
}

func (l *ChildList) MoveToFront(e *Child) {
	if e.list != l || l.front == e {
		return
	}
	l.remove(e)
	l.insertFront(e)
}

func (l *ChildList) MoveToBack(e *Child) {
	if e.list != l || l.back == e {
		return
	}
	l.remove(e)
	l.insertBack(e)
}

func (l *ChildList) Len() int { return l.len }

func (l *ChildList) PushBackList(ol *ChildList) {
	last := ol.Back()
	for e := ol.Front(); e != nil; e = e.Next() {
		l.PushBack(e.Node)
		if e == last {
			break
		}
	}
}

func (l *ChildList) PushFrontList(ol *ChildList) {
	first := ol.Front()
	for e := ol.Back(); e != nil; e = e.Prev() {
		l.PushFront(e.Node)
		if e == first {
			break
		}
	}
}
