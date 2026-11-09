package lib

import (
	"math/bits"
)

type Segtree[S any] struct {
	_data []S
	_n    int
	_size int
	_op   func(S, S) S
	_e    S
}

func NewSegtree[S any](data []S, N int, op func(S, S) S, e S) *Segtree[S] {
	_size := 1 << (bits.Len(uint(N - 1)))
	_data := make([]S, 2*_size)
	for i := range _data {
		_data[i] = e
	}
	for i := 0; i < N; i++ {
		_data[_size+i] = data[i]
	}
	seg := &Segtree[S]{
		_data: _data,
		_n:    N,
		_size: _size,
		_op:   op,
		_e:    e,
	}
	seg.update()

	return seg
}

func (seg *Segtree[S]) _prod(i int) {
	seg._data[i] = seg._op(seg._data[2*i], seg._data[2*i+1])
}

func (seg *Segtree[S]) update() {
	for i := seg._size - 1; i >= 1; i-- {
		seg._prod(i)
	}
}

func (seg *Segtree[S]) Get(i int) S {
	return seg._data[seg._size+i]
}

func (seg *Segtree[S]) Prod(l int, r int) S {
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

func (seg *Segtree[S]) All() S {
	return seg._data[1]
}

func (seg *Segtree[S]) Set(i int, x S) {
	i += seg._size
	seg._data[i] = x
	i >>= 1
	for i != 0 {
		seg._prod(i)
		i >>= 1
	}
}

func (seg *Segtree[S]) MaxRight(l int, f func(S) bool) int {
	if l == seg._n {
		return seg._n
	}
	l += seg._size
	sum := seg._e
	for {
		for l&1 == 0 {
			l >>= 1
		}
		if !f(seg._op(sum, seg._data[l])) {
			for l < seg._size {
				l <<= 1
				if f(seg._op(sum, seg._data[l])) {
					sum = seg._op(sum, seg._data[l])
					l++
				}
			}
			return l - seg._size
		}
		sum = seg._op(sum, seg._data[l])
		l++
		if l&(l-1) == 0 {
			break
		}
	}
	return seg._n
}

func (seg *Segtree[S]) MinLeft(r int, f func(S) bool) int {
	if r == 0 {
		return 0
	}
	r += seg._size
	sum := seg._e
	for {
		r--
		for r&1 != 0 && r > 1 {
			r >>= 1
		}
		if !f(seg._op(seg._data[r], sum)) {
			for r < seg._size {
				r <<= 1
				r++
				if f(seg._op(seg._data[r], sum)) {
					sum = seg._op(seg._data[r], sum)
					r--
				}
			}
			return r + 1 - seg._size
		}
		sum = seg._op(seg._data[r], sum)
		if r&(r-1) == 0 {
			break
		}
	}
	return 0
}
