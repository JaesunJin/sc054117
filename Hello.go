package main
import (
  "fmt"
  "strconv"
)

func add (x, y int) int {
  return x+y
}

func swap(x, y string) (string, string) {
  return y, x
}

func main() {
  s := strconv.Itoa(add(10,2))
  fmt.Printf("func %s test", s)

  fmt.Println()

  x, y := swap("world", "hello")
  fmt.Println(x, y)

}
