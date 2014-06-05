package output

import (
  "fmt"
  "jeremypholland.com/perco/percolation"
)

func RenderPerco(perco *percolation.Perco) {
  printTopBottom(perco)
  printRows(perco)
  printTopBottom(perco)
}

func printTopBottom(perco *percolation.Perco) {
  for i := 0; i < perco.Dim*2 + 3; i++ {
    fmt.Printf("=")
  }
  fmt.Printf("\n")
}

func printRows(perco *percolation.Perco) {
  for i := 0; i < perco.Dim; i++ {
    fmt.Printf("| ")
    for j := 0; j < perco.Dim; j++ {
      if perco.IsOpen(i, j) {
        if perco.IsFull(i, j) {
          fmt.Printf("\033[36m\u25A0\033[39m ")
        } else {
          fmt.Printf("\u25A0 ")
        }
      } else {
        fmt.Printf("  ")
      }
    }
    fmt.Printf("|\n")
  }
}
