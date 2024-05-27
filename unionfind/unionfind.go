package unionfind

type UnionFind []int

func NewUnionFind(n int) UnionFind {
	unionFind := make([]int, n)
	for i := range unionFind {
		unionFind[i] = i
	}
	return unionFind
}

func (uf UnionFind) Find(x int) int {
	for uf[x] != x {
		x, uf[x] = uf[x], uf[uf[x]]
	}
	return x
}

func (uf UnionFind) Join(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	uf[rootX] = rootY
}
