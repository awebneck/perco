package uf

type UnionFind struct {
  Ids []int
  Sizes []int
  Count int
}

func GenerateUnionFind(c int) (*UnionFind) {
  uf := new(UnionFind)
  uf.Count = c
  uf.Ids = make([]int, c + 2)
  for i := 0; i < c + 2; i++ {
    uf.Ids[i] = i
  }
  uf.Sizes = make([]int, c + 2)
  for i := 0; i < c + 2; i++ {
    uf.Sizes[i] = 1
  }
  return uf
}

func (uf *UnionFind) Union(i, j int) {
  idi := uf.Find(i)
  idj := uf.Find(j)
  if idi != idj {
    if uf.Sizes[idi] < uf.Sizes[idj] {
      uf.Ids[idi] = idj
      uf.Sizes[idj] += uf.Sizes[idi]
    } else {
      uf.Ids[idj] = idi
      uf.Sizes[idi] += uf.Sizes[idj]
    }
  }
}

func (uf *UnionFind) Find(i int) (int) {
  for i != uf.Ids[i] {
    i = uf.Ids[i]
  }
  return i
}

func (uf *UnionFind) Connected(i, j int) (bool) {
  return uf.Find(i) == uf.Find(j)
}
