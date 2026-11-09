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
	dq.last = (dq.last + 1) % len(dq.buf)
	dq.size++
}

func (dq *Deque[T]) PushLeft(val T) {
	if dq.size == len(dq.buf) {
		dq.grow()
	}
	dq.first = (dq.first - 1 + len(dq.buf)) % len(dq.buf)
	dq.buf[dq.first] = val
	dq.size++
}

func (dq *Deque[T]) Pop() T {
	if dq.size == 0 {
		panic(errors.New("Nothing to pop"))
	}
	dq.last = (dq.last - 1 + len(dq.buf)) % len(dq.buf)
	dq.size--
	return dq.buf[dq.last]
}

func (dq *Deque[T]) PopLeft() T {
	if dq.size == 0 {
		panic(errors.New("Nothing to pop"))
	}
	val := dq.buf[dq.first]
	dq.first = (dq.first + 1) % len(dq.buf)
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
	return dq.buf[(dq.first+i)%len(dq.buf)]
}
