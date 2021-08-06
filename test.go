package main

import (
  "fmt"
)

func main() {
  nums := []int{1, 2, 3, 4, 5}
  sum := 0
  for i, n := range nums {
    i = 6
    sum += n
  }
  fmt.Println(sum)
}