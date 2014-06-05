package percolation

import (
  "math/rand"
  "time"
  "jeremypholland.com/perco/uf"
)

type Perco struct {
  Grid [][]bool
  Dim int
  Density float64
  Uf *uf.UnionFind
}

func GeneratePerco(dim int, density float64) (*Perco, error) {
  rand.Seed(time.Now().UnixNano())
  perco := new(Perco)
  perco.Dim = dim
  perco.Density = density
  perco.Grid = make([][]bool, dim)
  perco.Uf = uf.GenerateUnionFind(dim*dim)
  for i := 0; i < dim; i++ {
    perco.Grid[i] = make([]bool, dim)
    for j := 0; j < dim; j++ {
      val := rand.Float64()
      perco.Grid[i][j] = val < density
    }
  }
  perco.unionize()
  return perco, nil
}

func (perco *Perco) IsOpen(i, j int) (bool) {
  return perco.Grid[i][j]
}

func (perco *Perco) IsFull(i, j int) (bool) {
  return perco.Uf.Connected(i*perco.Dim + j + 2, 0)
}

func (perco *Perco) Percolates() (bool) {
  return perco.Uf.Connected(1, 0)
}

func (perco *Perco) unionize() {
  for i := 0; i < perco.Dim; i++ {
    if perco.IsOpen(0, i) {
      perco.Uf.Union(0, i + 2)
    }
  }
  for i := 0; i < perco.Dim; i++ {
    if perco.IsOpen(perco.Dim - 1, i) {
      perco.Uf.Union(1, (perco.Dim - 1)*perco.Dim + i + 2)
    }
  }
  for i := 0; i < perco.Dim; i++ {
    for j := 0; j < perco.Dim; j++ {
      if perco.IsOpen(i, j) {
        if j < perco.Dim - 1 && perco.IsOpen(i, j + 1) {
          perco.Uf.Union(i*perco.Dim + j + 2, i*perco.Dim + j + 3)
        }
        if i < perco.Dim - 1 && perco.IsOpen(i + 1, j) {
          perco.Uf.Union(i*perco.Dim + j + 2, (i + 1)*perco.Dim + j + 2)
        }
      }
    }
  }
}
