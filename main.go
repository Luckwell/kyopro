package main

import (
	"kyopro/lib"
	"os"
)

type S struct {
	first  int
	second int
}

func main() {
	op := func(a S, b S) S {
		return S{
			first:  a.first + b.first,
			second: a.second + b.second,
		}
	}
	var N, Q int
	fastin := lib.NewScanner(os.Stdin)
	fastout := lib.NewWriter(os.Stdout)
	defer fastout.Flush()
	N = fastin.Int()
	Q = fastin.Int()
	A := make([]int, N)
	v := make([]S, 5e5+5)
	for i := range A {
		A[i] = fastin.Int()
		v[A[i]].first++
		v[A[i]].second += A[i]
	}
	seg := lib.NewSegtree(v, 5e5+5, op, S{first: 0, second: 0})
	for Q > 0 {
		Q--
		var t, x, y int
		t = fastin.Int()
		x = fastin.Int()
		y = fastin.Int()
		if t == 1 {
			x--
			tmp := seg.Get(A[x])
			tmp.first--
			tmp.second -= A[x]
			seg.Set(A[x], tmp)
			tmp = seg.Get(y)
			tmp.first++
			tmp.second += y
			seg.Set(y, tmp)
			A[x] = y
		} else {
			ans := 0
			if x < y {
				ans += x * seg.Prod(0, x).first
				ans += seg.Prod(x, y).second
				ans += seg.Prod(y, 5e5+1).first * y
			} else {
				ans += x * seg.All().first
			}
			fastout.PrintInt(ans)
			fastout.PrintByte('\n')
		}
	}
}
