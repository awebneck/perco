package main

import (
  "os"
  "fmt"
  "strconv"
  "errors"
  "github.com/awebneck/perco/percolation"
  "github.com/awebneck/perco/output"
)

func main() {
  dim, density, err := parseArgs()
  if err != nil {
    fmt.Printf(err.Error())
    os.Exit(1)
  }
  fmt.Printf("Generating sample 2D percolation of %dx%d with density %f\n", dim, dim, density)
  fmt.Printf("\n")
  perco, err := percolation.GeneratePerco(dim, density);
  output.RenderPerco(perco)
  if (perco.Percolates()) {
    fmt.Printf("PERCOLATES!!!\n")
  } else {
    fmt.Printf("DOES NOT PERCOLATE!!!\n")
  }
  if err != nil {
    fmt.Printf(err.Error())
    os.Exit(1)
  }
}

func parseArgs() (int, float64, error) {
  dim, err := strconv.Atoi(os.Args[1])
  if err != nil {
    return 0, 0.0, err
  }
  density := 0.5
  if len(os.Args) > 2 {
    density, err = strconv.ParseFloat(os.Args[2], 64)
    if err != nil {
      return 0, 0.0, err
    }
    if density >= 1.0 {
      return 0, 0.0, errors.New("Supplied density must be less than 1\n")
    }
  }
  return dim, density, nil
}
