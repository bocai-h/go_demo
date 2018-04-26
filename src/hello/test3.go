package main

import (
  "os"
  "log"
)

var cwd string

func init() {
  var err error
  cwd, err = os.Getwd()
  if err != nil {
    log.Fatalf("os,Getwd failed: %v", err)
  }
  log.Printf("Working directory = %s", cwd)
}

func main() {
  log.Printf("this is the main func")
}
