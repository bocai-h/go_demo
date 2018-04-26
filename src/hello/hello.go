package main

import "fmt"

func main() {
  fmt.Println("Hello, 世界", "world")
  fmt.Printf("20 24的最大公约数是%d\n", gcd(20, 24))
  fmt.Printf("第10个非波拉起数列是%d\n", fib(10))
}

func gcd(x,y int) int{
  for y != 0 {
    x, y = y, x % y
  }

  return x
}

func fib(n int) int {
  x, y := 0, 1
  for i := 0; i < n; i++ {
     x, y = y, x + y
  }
 return x
}
