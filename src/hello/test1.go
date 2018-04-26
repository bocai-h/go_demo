package main

import "fmt"

func main() {
  x := "hello!"
  fmt.Printf("a=%d\n", 'a')
  fmt.Printf("A=%d\n", 'A')
  for i := 0; i < len(x); i++ {
     x := x[i]
     if x != '!' {
         //x := x + 'A' - 'a'
         fmt.Printf("%c=%[1]d\n", x)
     }
  }
}
