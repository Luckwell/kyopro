package lib

import (
	"errors"
	"math/bits"
)

type Deque[T any] struct {
	buf   []T
	first int
	last  int
	size  int
}

func NewDeque[T any](capacity ...int) *Deque[T] {
	cap := 64
	if len(capacity) > 0 {
		cap = 1 << (bits.Len(uint(max(1, capacity[0]) - 1)))
	}
	return &Deque[T]{
		make([]T, cap),
		0,
		0,
		0,
	}
}

func (dq *Deque[T]) Push(val T) {
	if dq.size == len(dq.buf) {
		dq.grow()
	}
	dq.buf[dq.last] = val
	dq.last = (dq.last + 1) & (len(dq.buf) - 1)
	dq.size++
}

func (dq *Deque[T]) PushLeft(val T) {
	if dq.size == len(dq.buf) {
		dq.grow()
	}
	dq.first = (dq.first - 1) & (len(dq.buf) - 1)
	dq.buf[dq.first] = val
	dq.size++
}

func (dq *Deque[T]) Pop() T {
	if dq.size == 0 {
		panic(errors.New("Nothing to pop"))
	}
	dq.last = (dq.last - 1) & (len(dq.buf) - 1)
	dq.size--
	return dq.buf[dq.last]
}

func (dq *Deque[T]) PopLeft() T {
	if dq.size == 0 {
		panic(errors.New("Nothing to pop"))
	}
	val := dq.buf[dq.first]
	dq.first = (dq.first + 1) & (len(dq.buf) - 1)
	dq.size--
	return val
}

func (dq *Deque[T]) grow() {
	newCap := len(dq.buf) * 2
	newBuf := make([]T, newCap)

	n := copy(newBuf, dq.buf[dq.first:])
	copy(newBuf[n:], dq.buf[:dq.last])
	dq.buf = newBuf
	dq.first = 0
	dq.last = dq.size
}

func (dq *Deque[T]) Get(i int) T {
	if i < 0 || i >= dq.size {
		panic(errors.New("Index out of range"))
	}
	return dq.buf[(dq.first+i)&(len(dq.buf)-1)]
}

// minimum heap
type Lesser[T any] interface {
	Less(other T) bool
}

type PriorityQueue[T Lesser[T]] struct {
	_data []T
}

func NewPriorityQueue[T Lesser[T]](N ...int) *PriorityQueue[T] {
	size := 64
	if len(N) != 0 {
		size = max(size, N[0])
	}
	return &PriorityQueue[T]{
		make([]T, 0, size),
	}
}

func (pq *PriorityQueue[T]) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !pq._data[j].Less(pq._data[i]) {
			break
		}
		pq.swap(i, j)
		j = i
	}
}

func (pq *PriorityQueue[T]) down(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && pq._data[j2].Less(pq._data[j1]) {
			j = j2
		}
		if !pq._data[j].Less(pq._data[i]) {
			break
		}
		pq.swap(i, j)
		i = j
	}
}

func (pq *PriorityQueue[T]) swap(i, j int) {
	pq._data[i], pq._data[j] = pq._data[j], pq._data[i]
}

func (pq *PriorityQueue[T]) Push(elem T) {
	pq._data = append(pq._data, elem)
	pq.up(len(pq._data) - 1)
}

func (pq *PriorityQueue[T]) Pop() T {
	if len(pq._data) == 0 {
		panic("You are about to pop from nothing!")
	}
	old := pq._data
	n := len(old) - 1
	pq.swap(0, n)
	v := old[n]
	pq._data = old[0:n]
	pq.down(0, n)
	return v
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq._data)
}

func (pq *PriorityQueue[T]) Top() T {
	if len(pq._data) != 0 {
		return pq._data[0]
	} else {
		panic("No element inside this priority queue")
	}
}
