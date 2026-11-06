package lib

type dsu struct {
	_parent []int
}

func NewDSU(N int) *dsu {
	par := make([]int, N)
	for i := range par {
		par[i] = -1
	}
	return &dsu{
		_parent: par,
	}
}

func (u *dsu) Merge(a int, b int) int {
	a = u.Leader(a)
	b = u.Leader(b)
	if a == b {
		return a
	} else {
		if u._parent[a] < u._parent[b] {
			u._parent[a] += u._parent[b]
			u._parent[b] = a
			return a
		} else {
			u._parent[b] += u._parent[a]
			u._parent[a] = b
			return b
		}
	}
}

func (u *dsu) Leader(a int) int {
	if u._parent[a] < 0 {
		return a
	} else {
		u._parent[a] = u.Leader(u._parent[a])
		return u._parent[a]
	}
}

func (u *dsu) Same(a int, b int) bool {
	return u.Leader(a) == u.Leader(b)
}

func (u *dsu) Size(a int) int {
	return -u._parent[u.Leader(a)]
}
