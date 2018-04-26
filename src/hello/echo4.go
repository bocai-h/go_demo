package main
import(
  "fmt"
  "os"
)

func main(){
//  s, sep := "", ""
  fmt.Printf("\n")
  fmt.Print(os.Args[1:])
  fmt.Printf("\n")
  for index, arg := range os.Args[1:]{
    //s += sep + arg
    //sep = " "
    fmt.Println("index=", index, " value=", arg)
  }

//  fmt.Println(s)
}
