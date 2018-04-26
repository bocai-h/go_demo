package main

import "fmt"

func main() {
   str := "Hello,世界"
   for i, w := range str {
      fmt.Printf("%d\t%q\t%d\n", i, w, w)
   }
}
