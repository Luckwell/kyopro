package lib

import (
	"fmt"
	"math/rand/v2"
)

type TNode[T Lesser[T]] struct {
	_size       int
	value       T
	priority    int
	left, right *TNode[T]
}

func NewTNode[T Lesser[T]](value T) *TNode[T] {
	return &TNode[T]{
		1,
		value,
		rand.IntN(1 << 32),
		nil, nil,
	}
}
func (n *TNode[T]) size() int {
	if n != nil {
		return n._size
	}
	return 0
}

func (n *TNode[T]) update() {
	if n != nil {
		n._size = n.left.size() + n.right.size() + 1
	}
}
func merge[T Lesser[T]](l *TNode[T], r *TNode[T]) *TNode[T] {
	if l == nil || r == nil {
		if l == nil {
			return r
		} else {
			return l
		}
	}
	if l.priority > r.priority {
		l.right = merge(l.right, r)
		l.update()
		return l
	} else {
		r.left = merge(l, r.left)
		r.update()
		return r
	}
}

func split[T Lesser[T]](t *TNode[T], value T) (*TNode[T], *TNode[T]) {
	if t == nil {
		return nil, nil
	}
	if t.value.Less(value) {
		merg, spl := split(t.right, value)
		t.right = merg
		t.update()
		return t, spl
	} else {
		spl, merg := split(t.left, value)
		t.left = merg
		t.update()
		return spl, t
	}
}

func spliteq[T Lesser[T]](t *TNode[T], value T) (*TNode[T], *TNode[T]) {
	if t == nil {
		return nil, nil
	}
	if !value.Less(t.value) {
		merge, spl := spliteq(t.right, value)
		t.right = merge
		t.update()
		return t, spl
	} else {
		spl, merg := spliteq(t.left, value)
		t.left = merg
		t.update()
		return spl, t
	}
}

type Treap[T Lesser[T]] struct {
	_data []TNode[T]
	_root *TNode[T]
	cap   int
}

func NewTreap[T Lesser[T]](cap int) *Treap[T] {
	return &Treap[T]{
		make([]TNode[T], 0, cap),
		nil,
		cap,
	}
}

func (t *Treap[T]) Insert(elem T) {
	left, right := split(t._root, elem)
	exist, right := spliteq(right, elem)
	if exist != nil {
		t._root = merge(left, merge(exist, right))
		return
	}
	if len(t._data) == t.cap {
		panic(fmt.Sprintf("You exceeded initialized treap cap of %d", t.cap))
	}
	t._data = append(t._data, *NewTNode(elem))
	t._root = merge(merge(left, &t._data[len(t._data)-1]), right)
}

func (t *Treap[T]) Delete(elem T) {
	left, mid_right := split(t._root, elem)
	_, right := spliteq(mid_right, elem)
	t._root = merge(left, right)
}

func (t *Treap[T]) Contains(elem T) bool {
	cur := t._root
	for {
		if cur == nil {
			return false
		}
		tmp := cur.value.Less(elem)
		if !tmp && !elem.Less(cur.value) {
			return true
		}
		if tmp {
			cur = cur.right
		} else {
			cur = cur.left
		}
	}
}

func (t *Treap[T]) At(ind int) T {
	if ind >= t._root.size() || ind < 0 {
		panic("Index out of range")
	}
	cur := t._root
	for {
		ls := cur.left.size()
		if ls < ind {
			ind -= ls + 1
			cur = cur.right
		} else if ls == ind {
			return cur.value
		} else {
			cur = cur.left
		}
	}
}
