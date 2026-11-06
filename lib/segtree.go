package lib

import (
	"math/bits"
)

type Segtree[T any] struct {
	_data []T
	_n    int
	_size int
	_op   func(T, T) T
	_e    T
}

func NewSegtree[T any](data []T, N int, op func(T, T) T, e T) *Segtree[T] {
	_size := 1 << (bits.Len(uint(N - 1)))
	_data := make([]T, 2*_size)
	for i := range _data {
		_data[i] = e
	}
	for i := 0; i < N; i++ {
		_data[_size+i] = data[i]
	}
	seg := &Segtree[T]{
		_data: _data,
		_n:    N,
		_size: _size,
		_op:   op,
		_e:    e,
	}
	seg.update()

	return seg
}

func (seg *Segtree[T]) _prod(i int) {
	seg._data[i] = seg._op(seg._data[2*i], seg._data[2*i+1])
}

func (seg *Segtree[T]) update() {
	for i := seg._size - 1; i >= 1; i-- {
		seg._prod(i)
	}
}

func (seg *Segtree[T]) Get(i int) T {
	return seg._data[seg._size+i]
}

func (seg *Segtree[T]) Prod(l int, r int) T {
	l += seg._size
	r += seg._size
	opl := seg._e
	opr := seg._e
	for l < r {
		if l&1 != 0 {
			opl = seg._op(opl, seg._data[l])
			l++
		}
		if r&1 != 0 {
			r--
			opr = seg._op(seg._data[r], opr)
		}
		l >>= 1
		r >>= 1
	}
	return seg._op(opl, opr)
}

func (seg *Segtree[T]) All() T {
	return seg._data[1]
}

func (seg *Segtree[T]) Set(i int, x T) {
	i += seg._size
	seg._data[i] = x
	i >>= 1
	for i != 0 {
		seg._prod(i)
		i >>= 1
	}
}

func (seg *Segtree[T]) MaxRight(l int, f func()) {

}
