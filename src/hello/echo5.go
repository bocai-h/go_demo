// Echo5 prints its command-line arguments

package main

import (
  "flag"
  "fmt"
  "strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main(){
  flag.Parse()
  fmt.Print(strings.Join(flag.Args(), *sep))
  fmt.Printf("*n=")
  fmt.Print(*n)
  if !*n {
    fmt.Println()
  }
}
